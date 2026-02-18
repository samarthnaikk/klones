package session

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// Session represents a playback session
type Session struct {
	SessionID    string    `json:"session_id"`
	UserID       string    `json:"user_id"`
	ProfileID    string    `json:"profile_id"`
	ContentID    string    `json:"content_id"`
	DeviceID     string    `json:"device_id"`
	Status       string    `json:"status"` // active, paused, stopped
	StartedAt    time.Time `json:"started_at"`
	LastHeartbeat time.Time `json:"last_heartbeat"`
	Position     int64     `json:"position"` // playback position in seconds
}

// Manager handles session operations
type Manager struct {
	redisClient *redis.Client
	timeout     time.Duration
}

// NewManager creates a new session manager
func NewManager(redisClient *redis.Client, timeout time.Duration) *Manager {
	return &Manager{
		redisClient: redisClient,
		timeout:     timeout,
	}
}

// CreateSession creates a new playback session
func (m *Manager) CreateSession(ctx context.Context, userID, profileID, contentID, deviceID string) (*Session, error) {
	session := &Session{
		SessionID:     uuid.New().String(),
		UserID:        userID,
		ProfileID:     profileID,
		ContentID:     contentID,
		DeviceID:      deviceID,
		Status:        "active",
		StartedAt:     time.Now(),
		LastHeartbeat: time.Now(),
		Position:      0,
	}

	// Store session in Redis
	key := m.getSessionKey(session.SessionID)
	data, err := json.Marshal(session)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal session: %w", err)
	}

	err = m.redisClient.Set(ctx, key, data, m.timeout).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to store session in Redis: %w", err)
	}

	// Track active session by user
	userSessionKey := m.getUserSessionKey(userID)
	err = m.redisClient.SAdd(ctx, userSessionKey, session.SessionID).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to track user session: %w", err)
	}

	return session, nil
}

// GetSession retrieves a session by ID
func (m *Manager) GetSession(ctx context.Context, sessionID string) (*Session, error) {
	key := m.getSessionKey(sessionID)
	data, err := m.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("session not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve session: %w", err)
	}

	var session Session
	err = json.Unmarshal([]byte(data), &session)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	return &session, nil
}

// UpdateHeartbeat updates the last heartbeat time and position
func (m *Manager) UpdateHeartbeat(ctx context.Context, sessionID string, position int64) error {
	session, err := m.GetSession(ctx, sessionID)
	if err != nil {
		return err
	}

	session.LastHeartbeat = time.Now()
	session.Position = position

	key := m.getSessionKey(sessionID)
	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	err = m.redisClient.Set(ctx, key, data, m.timeout).Err()
	if err != nil {
		return fmt.Errorf("failed to update session: %w", err)
	}

	return nil
}

// UpdateStatus updates session status
func (m *Manager) UpdateStatus(ctx context.Context, sessionID, status string) error {
	session, err := m.GetSession(ctx, sessionID)
	if err != nil {
		return err
	}

	session.Status = status

	key := m.getSessionKey(sessionID)
	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	err = m.redisClient.Set(ctx, key, data, m.timeout).Err()
	if err != nil {
		return fmt.Errorf("failed to update session: %w", err)
	}

	return nil
}

// TerminateSession terminates a session
func (m *Manager) TerminateSession(ctx context.Context, sessionID string) error {
	session, err := m.GetSession(ctx, sessionID)
	if err != nil {
		return err
	}

	// Remove from active sessions
	userSessionKey := m.getUserSessionKey(session.UserID)
	err = m.redisClient.SRem(ctx, userSessionKey, sessionID).Err()
	if err != nil {
		return fmt.Errorf("failed to remove from user sessions: %w", err)
	}

	// Delete session
	key := m.getSessionKey(sessionID)
	err = m.redisClient.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	return nil
}

// GetActiveSessions returns all active sessions for a user
func (m *Manager) GetActiveSessions(ctx context.Context, userID string) ([]string, error) {
	userSessionKey := m.getUserSessionKey(userID)
	sessionIDs, err := m.redisClient.SMembers(ctx, userSessionKey).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user sessions: %w", err)
	}

	return sessionIDs, nil
}

// Helper functions for Redis keys
func (m *Manager) getSessionKey(sessionID string) string {
	return fmt.Sprintf("session:%s", sessionID)
}

func (m *Manager) getUserSessionKey(userID string) string {
	return fmt.Sprintf("user:%s:sessions", userID)
}
