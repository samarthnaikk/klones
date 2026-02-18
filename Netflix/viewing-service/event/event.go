package event

import (
	"encoding/json"
	"log"
	"time"
)

// EventType represents different types of playback events
type EventType string

const (
	EventPlaybackStarted   EventType = "playback.started"
	EventPlaybackPaused    EventType = "playback.paused"
	EventPlaybackResumed   EventType = "playback.resumed"
	EventPlaybackStopped   EventType = "playback.stopped"
	EventPlaybackCompleted EventType = "playback.completed"
	EventHeartbeat         EventType = "playback.heartbeat"
)

// Event represents a playback event
type Event struct {
	EventType EventType              `json:"event_type"`
	SessionID string                 `json:"session_id"`
	UserID    string                 `json:"user_id"`
	ProfileID string                 `json:"profile_id"`
	ContentID string                 `json:"content_id"`
	DeviceID  string                 `json:"device_id"`
	Timestamp time.Time              `json:"timestamp"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// Emitter handles event emission
type Emitter struct {
	// In a production system, this would integrate with a message queue
	// (Kafka, RabbitMQ, etc.) or event bus for the Engagement service
}

// NewEmitter creates a new event emitter
func NewEmitter() *Emitter {
	return &Emitter{}
}

// EmitPlaybackStarted emits a playback started event
func (e *Emitter) EmitPlaybackStarted(sessionID, userID, profileID, contentID, deviceID string, metadata map[string]interface{}) {
	event := Event{
		EventType: EventPlaybackStarted,
		SessionID: sessionID,
		UserID:    userID,
		ProfileID: profileID,
		ContentID: contentID,
		DeviceID:  deviceID,
		Timestamp: time.Now(),
		Metadata:  metadata,
	}
	e.emit(event)
}

// EmitPlaybackPaused emits a playback paused event
func (e *Emitter) EmitPlaybackPaused(sessionID, userID, profileID, contentID, deviceID string, position int64) {
	event := Event{
		EventType: EventPlaybackPaused,
		SessionID: sessionID,
		UserID:    userID,
		ProfileID: profileID,
		ContentID: contentID,
		DeviceID:  deviceID,
		Timestamp: time.Now(),
		Metadata: map[string]interface{}{
			"position": position,
		},
	}
	e.emit(event)
}

// EmitPlaybackResumed emits a playback resumed event
func (e *Emitter) EmitPlaybackResumed(sessionID, userID, profileID, contentID, deviceID string, position int64) {
	event := Event{
		EventType: EventPlaybackResumed,
		SessionID: sessionID,
		UserID:    userID,
		ProfileID: profileID,
		ContentID: contentID,
		DeviceID:  deviceID,
		Timestamp: time.Now(),
		Metadata: map[string]interface{}{
			"position": position,
		},
	}
	e.emit(event)
}

// EmitPlaybackStopped emits a playback stopped event
func (e *Emitter) EmitPlaybackStopped(sessionID, userID, profileID, contentID, deviceID string, position int64) {
	event := Event{
		EventType: EventPlaybackStopped,
		SessionID: sessionID,
		UserID:    userID,
		ProfileID: profileID,
		ContentID: contentID,
		DeviceID:  deviceID,
		Timestamp: time.Now(),
		Metadata: map[string]interface{}{
			"position": position,
		},
	}
	e.emit(event)
}

// EmitPlaybackCompleted emits a playback completed event
func (e *Emitter) EmitPlaybackCompleted(sessionID, userID, profileID, contentID, deviceID string) {
	event := Event{
		EventType: EventPlaybackCompleted,
		SessionID: sessionID,
		UserID:    userID,
		ProfileID: profileID,
		ContentID: contentID,
		DeviceID:  deviceID,
		Timestamp: time.Now(),
	}
	e.emit(event)
}

// EmitHeartbeat emits a heartbeat event
func (e *Emitter) EmitHeartbeat(sessionID, userID, profileID, contentID, deviceID string, position int64) {
	event := Event{
		EventType: EventHeartbeat,
		SessionID: sessionID,
		UserID:    userID,
		ProfileID: profileID,
		ContentID: contentID,
		DeviceID:  deviceID,
		Timestamp: time.Now(),
		Metadata: map[string]interface{}{
			"position": position,
		},
	}
	e.emit(event)
}

// emit sends an event to the event pipeline
func (e *Emitter) emit(event Event) {
	// In production, this would publish to a message queue
	// For now, we'll log the event
	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshaling event: %v", err)
		return
	}

	log.Printf("Event emitted: %s", eventJSON)
	
	// TODO: Integrate with message queue (Kafka, RabbitMQ, etc.)
	// Example:
	// - Publish to Kafka topic
	// - Send to RabbitMQ exchange
	// - Push to Redis pub/sub
	// - Send to cloud event bus (AWS EventBridge, GCP Pub/Sub, etc.)
}
