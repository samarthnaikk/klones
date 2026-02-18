package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/samarthnaikk/klones/Netflix/viewing-service/concurrency"
	"github.com/samarthnaikk/klones/Netflix/viewing-service/config"
	"github.com/samarthnaikk/klones/Netflix/viewing-service/entitlement"
	"github.com/samarthnaikk/klones/Netflix/viewing-service/event"
	"github.com/samarthnaikk/klones/Netflix/viewing-service/lifecycle"
	"github.com/samarthnaikk/klones/Netflix/viewing-service/session"
	"github.com/samarthnaikk/klones/Netflix/viewing-service/token"
)

// Service holds all dependencies for the Viewing Service
type Service struct {
	config             *config.Config
	redisClient        *redis.Client
	sessionManager     *session.Manager
	entitlementValidator *entitlement.Validator
	concurrencyManager *concurrency.Manager
	tokenGenerator     *token.Generator
	lifecycleManager   *lifecycle.Manager
	eventEmitter       *event.Emitter
}

func main() {
	log.Println("Starting Viewing Service...")

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	// Test Redis connection
	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Printf("Warning: Redis connection failed: %v. Service will start but session storage will not work.", err)
	} else {
		log.Println("Redis connected successfully")
	}

	// Initialize service components
	svc := &Service{
		config:               cfg,
		redisClient:          redisClient,
		sessionManager:       session.NewManager(redisClient, cfg.SessionTimeout),
		entitlementValidator: entitlement.NewValidator(cfg.IdentityServiceURL),
		concurrencyManager:   concurrency.NewManager(redisClient, cfg.MaxConcurrentStreams),
		tokenGenerator:       token.NewGenerator(cfg.JWTSecret, cfg.JWTExpiration),
		lifecycleManager:     nil, // Will be set after session manager is created
		eventEmitter:         event.NewEmitter(),
	}

	// Initialize lifecycle manager with session manager
	svc.lifecycleManager = lifecycle.NewManager(svc.sessionManager, cfg.HeartbeatInterval)

	// Start heartbeat monitor in background
	go svc.lifecycleManager.StartHeartbeatMonitor(ctx)

	// Setup HTTP server
	router := svc.setupRoutes()
	server := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Viewing Service listening on port %s", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Close Redis connection
	if err := redisClient.Close(); err != nil {
		log.Printf("Error closing Redis connection: %v", err)
	}

	log.Println("Viewing Service stopped")
}

// setupRoutes configures all HTTP routes
func (s *Service) setupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Health check
	router.HandleFunc("/health", s.handleHealth).Methods("GET")

	// Playback session endpoints
	router.HandleFunc("/api/playback/session", s.handleCreateSession).Methods("POST")
	router.HandleFunc("/api/playback/session/{sessionID}", s.handleGetSession).Methods("GET")
	router.HandleFunc("/api/playback/session/{sessionID}/heartbeat", s.handleHeartbeat).Methods("POST")
	router.HandleFunc("/api/playback/session/{sessionID}/pause", s.handlePause).Methods("POST")
	router.HandleFunc("/api/playback/session/{sessionID}/resume", s.handleResume).Methods("POST")
	router.HandleFunc("/api/playback/session/{sessionID}/stop", s.handleStop).Methods("POST")
	router.HandleFunc("/api/playback/session/{sessionID}/terminate", s.handleTerminate).Methods("DELETE")

	// Token endpoints
	router.HandleFunc("/api/playback/token/validate", s.handleValidateToken).Methods("POST")
	router.HandleFunc("/api/playback/token/refresh", s.handleRefreshToken).Methods("POST")

	// Concurrency endpoints
	router.HandleFunc("/api/concurrency/{userID}", s.handleGetConcurrency).Methods("GET")

	return router
}

// Handler implementations

func (s *Service) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
		"service": "viewing-service",
	})
}

func (s *Service) handleCreateSession(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID    string `json:"user_id"`
		ProfileID string `json:"profile_id"`
		ContentID string `json:"content_id"`
		DeviceID  string `json:"device_id"`
		Region    string `json:"region,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request: %v", err), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	// Step 1: Validate entitlement
	entitled, err := s.entitlementValidator.ValidateEntitlement(ctx, req.UserID, req.ProfileID, req.ContentID)
	if err != nil || !entitled {
		s.respondError(w, fmt.Sprintf("Entitlement validation failed: %v", err), http.StatusForbidden)
		return
	}

	// Step 2: Validate content availability (if region is provided)
	if req.Region != "" {
		available, err := s.entitlementValidator.ValidateContentAvailability(ctx, req.ContentID, req.Region, s.config.ContentServiceURL)
		if err != nil || !available {
			s.respondError(w, fmt.Sprintf("Content not available: %v", err), http.StatusForbidden)
			return
		}
	}

	// Step 3: Create session
	sess, err := s.sessionManager.CreateSession(ctx, req.UserID, req.ProfileID, req.ContentID, req.DeviceID)
	if err != nil {
		s.respondError(w, fmt.Sprintf("Failed to create session: %v", err), http.StatusInternalServerError)
		return
	}

	// Step 4: Acquire concurrency slot
	acquired, err := s.concurrencyManager.AcquireSlot(ctx, req.UserID, sess.SessionID)
	if err != nil || !acquired {
		// Rollback session creation
		s.sessionManager.TerminateSession(ctx, sess.SessionID)
		s.respondError(w, fmt.Sprintf("Failed to acquire concurrency slot: %v", err), http.StatusTooManyRequests)
		return
	}

	// Step 5: Generate playback token
	playbackToken, err := s.tokenGenerator.GeneratePlaybackToken(sess.SessionID, req.UserID, req.ProfileID, req.ContentID, req.DeviceID)
	if err != nil {
		// Rollback
		s.concurrencyManager.ReleaseSlot(ctx, req.UserID, sess.SessionID)
		s.sessionManager.TerminateSession(ctx, sess.SessionID)
		s.respondError(w, fmt.Sprintf("Failed to generate token: %v", err), http.StatusInternalServerError)
		return
	}

	// Step 6: Emit playback started event
	s.eventEmitter.EmitPlaybackStarted(sess.SessionID, req.UserID, req.ProfileID, req.ContentID, req.DeviceID, nil)

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"session":        sess,
		"playback_token": playbackToken,
	})
}

func (s *Service) handleGetSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["sessionID"]

	sess, err := s.sessionManager.GetSession(r.Context(), sessionID)
	if err != nil {
		s.respondError(w, fmt.Sprintf("Session not found: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sess)
}

func (s *Service) handleHeartbeat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["sessionID"]

	var req struct {
		Position int64 `json:"position"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request: %v", err), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	// Get session to extract user info
	sess, err := s.sessionManager.GetSession(ctx, sessionID)
	if err != nil {
		s.respondError(w, fmt.Sprintf("Session not found: %v", err), http.StatusNotFound)
		return
	}

	if err := s.lifecycleManager.HandleHeartbeat(ctx, sessionID, req.Position); err != nil {
		s.respondError(w, fmt.Sprintf("Failed to update heartbeat: %v", err), http.StatusInternalServerError)
		return
	}

	// Emit heartbeat event
	s.eventEmitter.EmitHeartbeat(sess.SessionID, sess.UserID, sess.ProfileID, sess.ContentID, sess.DeviceID, req.Position)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (s *Service) handlePause(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["sessionID"]

	ctx := r.Context()

	// Get session to extract user info for event
	sess, err := s.sessionManager.GetSession(ctx, sessionID)
	if err != nil {
		s.respondError(w, fmt.Sprintf("Session not found: %v", err), http.StatusNotFound)
		return
	}

	if err := s.lifecycleManager.PauseSession(ctx, sessionID); err != nil {
		s.respondError(w, fmt.Sprintf("Failed to pause session: %v", err), http.StatusInternalServerError)
		return
	}

	// Emit pause event
	s.eventEmitter.EmitPlaybackPaused(sess.SessionID, sess.UserID, sess.ProfileID, sess.ContentID, sess.DeviceID, sess.Position)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "paused"})
}

func (s *Service) handleResume(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["sessionID"]

	ctx := r.Context()

	// Get session to extract user info for event
	sess, err := s.sessionManager.GetSession(ctx, sessionID)
	if err != nil {
		s.respondError(w, fmt.Sprintf("Session not found: %v", err), http.StatusNotFound)
		return
	}

	if err := s.lifecycleManager.ResumeSession(ctx, sessionID); err != nil {
		s.respondError(w, fmt.Sprintf("Failed to resume session: %v", err), http.StatusInternalServerError)
		return
	}

	// Emit resume event
	s.eventEmitter.EmitPlaybackResumed(sess.SessionID, sess.UserID, sess.ProfileID, sess.ContentID, sess.DeviceID, sess.Position)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "active"})
}

func (s *Service) handleStop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["sessionID"]

	ctx := r.Context()

	// Get session to extract user info for event
	sess, err := s.sessionManager.GetSession(ctx, sessionID)
	if err != nil {
		s.respondError(w, fmt.Sprintf("Session not found: %v", err), http.StatusNotFound)
		return
	}

	if err := s.lifecycleManager.StopSession(ctx, sessionID); err != nil {
		s.respondError(w, fmt.Sprintf("Failed to stop session: %v", err), http.StatusInternalServerError)
		return
	}

	// Emit stop event
	s.eventEmitter.EmitPlaybackStopped(sess.SessionID, sess.UserID, sess.ProfileID, sess.ContentID, sess.DeviceID, sess.Position)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "stopped"})
}

func (s *Service) handleTerminate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["sessionID"]

	ctx := r.Context()

	// Get session to extract user info before termination
	sess, err := s.sessionManager.GetSession(ctx, sessionID)
	if err != nil {
		s.respondError(w, fmt.Sprintf("Session not found: %v", err), http.StatusNotFound)
		return
	}

	// Release concurrency slot
	if err := s.concurrencyManager.ReleaseSlot(ctx, sess.UserID, sessionID); err != nil {
		log.Printf("Warning: Failed to release concurrency slot: %v", err)
	}

	// Terminate session
	if err := s.sessionManager.TerminateSession(ctx, sessionID); err != nil {
		s.respondError(w, fmt.Sprintf("Failed to terminate session: %v", err), http.StatusInternalServerError)
		return
	}

	// Emit completed event
	s.eventEmitter.EmitPlaybackCompleted(sess.SessionID, sess.UserID, sess.ProfileID, sess.ContentID, sess.DeviceID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "terminated"})
}

func (s *Service) handleValidateToken(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request: %v", err), http.StatusBadRequest)
		return
	}

	claims, err := s.tokenGenerator.ValidateToken(req.Token)
	if err != nil {
		s.respondError(w, fmt.Sprintf("Token validation failed: %v", err), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"valid":  true,
		"claims": claims,
	})
}

func (s *Service) handleRefreshToken(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request: %v", err), http.StatusBadRequest)
		return
	}

	newToken, err := s.tokenGenerator.RefreshToken(req.Token)
	if err != nil {
		s.respondError(w, fmt.Sprintf("Token refresh failed: %v", err), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": newToken,
	})
}

func (s *Service) handleGetConcurrency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	ctx := r.Context()

	count, err := s.concurrencyManager.GetActiveStreamCount(ctx, userID)
	if err != nil {
		s.respondError(w, fmt.Sprintf("Failed to get concurrency info: %v", err), http.StatusInternalServerError)
		return
	}

	sessions, err := s.concurrencyManager.GetActiveSessions(ctx, userID)
	if err != nil {
		s.respondError(w, fmt.Sprintf("Failed to get active sessions: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":                userID,
		"active_stream_count":    count,
		"max_concurrent_streams": s.config.MaxConcurrentStreams,
		"active_sessions":        sessions,
	})
}

// Helper function to send error responses
func (s *Service) respondError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
