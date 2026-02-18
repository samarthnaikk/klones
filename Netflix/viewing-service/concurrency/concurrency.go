package concurrency

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Manager handles concurrency enforcement
type Manager struct {
	redisClient          *redis.Client
	maxConcurrentStreams int
}

// NewManager creates a new concurrency manager
func NewManager(redisClient *redis.Client, maxConcurrentStreams int) *Manager {
	return &Manager{
		redisClient:          redisClient,
		maxConcurrentStreams: maxConcurrentStreams,
	}
}

// AcquireSlot attempts to acquire a concurrency slot for a user
func (m *Manager) AcquireSlot(ctx context.Context, userID, sessionID string) (bool, error) {
	lockKey := m.getConcurrencyKey(userID)

	// Use Redis SET with NX (only if not exists) and timeout for distributed locking
	// First, check current count
	count, err := m.redisClient.SCard(ctx, lockKey).Result()
	if err != nil && err != redis.Nil {
		return false, fmt.Errorf("failed to check concurrency count: %w", err)
	}

	if int(count) >= m.maxConcurrentStreams {
		return false, fmt.Errorf("maximum concurrent streams (%d) reached for user %s", m.maxConcurrentStreams, userID)
	}

	// Add session to the set of active streams
	err = m.redisClient.SAdd(ctx, lockKey, sessionID).Err()
	if err != nil {
		return false, fmt.Errorf("failed to acquire slot: %w", err)
	}

	// Set expiration on the key (safety mechanism)
	err = m.redisClient.Expire(ctx, lockKey, 24*time.Hour).Err()
	if err != nil {
		return false, fmt.Errorf("failed to set expiration: %w", err)
	}

	return true, nil
}

// ReleaseSlot releases a concurrency slot
func (m *Manager) ReleaseSlot(ctx context.Context, userID, sessionID string) error {
	lockKey := m.getConcurrencyKey(userID)

	err := m.redisClient.SRem(ctx, lockKey, sessionID).Err()
	if err != nil {
		return fmt.Errorf("failed to release slot: %w", err)
	}

	return nil
}

// GetActiveStreamCount returns the current number of active streams for a user
func (m *Manager) GetActiveStreamCount(ctx context.Context, userID string) (int, error) {
	lockKey := m.getConcurrencyKey(userID)

	count, err := m.redisClient.SCard(ctx, lockKey).Result()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, fmt.Errorf("failed to get active stream count: %w", err)
	}

	return int(count), nil
}

// GetActiveSessions returns the list of active session IDs for a user
func (m *Manager) GetActiveSessions(ctx context.Context, userID string) ([]string, error) {
	lockKey := m.getConcurrencyKey(userID)

	sessions, err := m.redisClient.SMembers(ctx, lockKey).Result()
	if err == redis.Nil {
		return []string{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get active sessions: %w", err)
	}

	return sessions, nil
}

// Helper function to get the concurrency key
func (m *Manager) getConcurrencyKey(userID string) string {
	return fmt.Sprintf("concurrency:%s", userID)
}
