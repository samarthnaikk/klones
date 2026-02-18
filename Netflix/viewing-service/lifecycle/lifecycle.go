package lifecycle

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/samarthnaikk/klones/Netflix/viewing-service/session"
)

// Manager handles session lifecycle management
type Manager struct {
	sessionManager    *session.Manager
	heartbeatInterval time.Duration
}

// NewManager creates a new lifecycle manager
func NewManager(sessionManager *session.Manager, heartbeatInterval time.Duration) *Manager {
	return &Manager{
		sessionManager:    sessionManager,
		heartbeatInterval: heartbeatInterval,
	}
}

// StartHeartbeatMonitor starts monitoring for stale sessions
func (m *Manager) StartHeartbeatMonitor(ctx context.Context) {
	ticker := time.NewTicker(m.heartbeatInterval)
	defer ticker.Stop()

	log.Println("Heartbeat monitor started")

	for {
		select {
		case <-ctx.Done():
			log.Println("Heartbeat monitor stopped")
			return
		case <-ticker.C:
			// This would scan for stale sessions in a production system
			// For now, we'll just log that the monitor is running
			log.Println("Heartbeat monitor tick - checking for stale sessions")
		}
	}
}

// HandleHeartbeat processes a heartbeat update
func (m *Manager) HandleHeartbeat(ctx context.Context, sessionID string, position int64) error {
	err := m.sessionManager.UpdateHeartbeat(ctx, sessionID, position)
	if err != nil {
		return fmt.Errorf("failed to update heartbeat: %w", err)
	}

	log.Printf("Heartbeat received for session %s at position %d", sessionID, position)
	return nil
}

// PauseSession pauses a playback session
func (m *Manager) PauseSession(ctx context.Context, sessionID string) error {
	err := m.sessionManager.UpdateStatus(ctx, sessionID, "paused")
	if err != nil {
		return fmt.Errorf("failed to pause session: %w", err)
	}

	log.Printf("Session %s paused", sessionID)
	return nil
}

// ResumeSession resumes a paused session
func (m *Manager) ResumeSession(ctx context.Context, sessionID string) error {
	err := m.sessionManager.UpdateStatus(ctx, sessionID, "active")
	if err != nil {
		return fmt.Errorf("failed to resume session: %w", err)
	}

	log.Printf("Session %s resumed", sessionID)
	return nil
}

// StopSession stops a playback session
func (m *Manager) StopSession(ctx context.Context, sessionID string) error {
	err := m.sessionManager.UpdateStatus(ctx, sessionID, "stopped")
	if err != nil {
		return fmt.Errorf("failed to stop session: %w", err)
	}

	log.Printf("Session %s stopped", sessionID)
	return nil
}

// GetSessionState retrieves the current state of a session
func (m *Manager) GetSessionState(ctx context.Context, sessionID string) (*session.Session, error) {
	sess, err := m.sessionManager.GetSession(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get session state: %w", err)
	}

	return sess, nil
}
