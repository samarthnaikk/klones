# Viewing Service

The Viewing Service is the control-plane coordinator of playback for the Netflix-like streaming platform.

## Overview

This service is responsible for:
- Creating playback sessions
- Validating entitlement with Identity Service
- Verifying content availability with Content Service
- Enforcing concurrency rules
- Issuing JWT-based playback tokens
- Managing playback lifecycle state
- Emitting playback events for Engagement Service

**Note:** This service does NOT stream video directly. It authorizes and orchestrates playback.

## Architecture

### Technology Stack
- **Language:** Go 1.24+
- **State Storage:** Redis (for session tracking, concurrency locks, heartbeats)
- **Optional:** PostgreSQL (for durable session logs)
- **Authentication:** JWT-based playback tokens
- **Design:** Stateless, horizontally scalable service

### Module Structure

```
viewing-service/
├── ViewingService.go      # Main entry point
├── config/                # Configuration management
├── session/               # Session creation, tracking, termination
├── entitlement/           # Entitlement validation (Identity Service integration)
├── concurrency/           # Slot acquisition, release, tracking
├── token/                 # JWT playback token generation
├── lifecycle/             # Heartbeat updates, session state management
└── event/                 # Event emission hooks (for Engagement)
```

## Configuration

All configuration is done via environment variables:

### Server Configuration
- `SERVER_PORT` - HTTP server port (default: 8080)

### Redis Configuration
- `REDIS_ADDR` - Redis address (default: localhost:6379)
- `REDIS_PASSWORD` - Redis password (default: "")
- `REDIS_DB` - Redis database number (default: 0)

### PostgreSQL Configuration (Optional)
- `POSTGRES_ENABLED` - Enable PostgreSQL (default: false)
- `POSTGRES_HOST` - PostgreSQL host (default: localhost)
- `POSTGRES_PORT` - PostgreSQL port (default: 5432)
- `POSTGRES_USER` - PostgreSQL user (default: postgres)
- `POSTGRES_PASSWORD` - PostgreSQL password
- `POSTGRES_DB` - PostgreSQL database name (default: viewing_service)

### JWT Configuration
- `JWT_SECRET` - Secret key for signing tokens (default: "default-secret-change-in-production")
- `JWT_EXPIRATION` - Token expiration duration (default: 1h)

### Service URLs
- `IDENTITY_SERVICE_URL` - Identity Service URL (default: http://localhost:8081)
- `CONTENT_SERVICE_URL` - Content Service URL (default: http://localhost:8082)

### Concurrency Configuration
- `MAX_CONCURRENT_STREAMS` - Maximum concurrent streams per account (default: 4)

### Session Configuration
- `SESSION_TIMEOUT` - Session timeout duration (default: 24h)
- `HEARTBEAT_INTERVAL` - Heartbeat monitoring interval (default: 30s)

## Building and Running

### Build
```bash
go build -o viewing-service ViewingService.go
```

### Run
```bash
./viewing-service
```

### Run with Custom Configuration
```bash
export SERVER_PORT=8080
export REDIS_ADDR=localhost:6379
export JWT_SECRET=my-secret-key
export MAX_CONCURRENT_STREAMS=4
./viewing-service
```

## API Endpoints

### Health Check
```
GET /health
```

### Playback Session Management

#### Create Session
```
POST /api/playback/session
Content-Type: application/json

{
  "user_id": "user123",
  "profile_id": "profile456",
  "content_id": "content789",
  "device_id": "device101",
  "region": "US"  // optional
}

Response:
{
  "session": {
    "session_id": "uuid",
    "user_id": "user123",
    "profile_id": "profile456",
    "content_id": "content789",
    "device_id": "device101",
    "status": "active",
    "started_at": "2026-02-18T08:54:22Z",
    "last_heartbeat": "2026-02-18T08:54:22Z",
    "position": 0
  },
  "playback_token": "jwt-token-here"
}
```

#### Get Session
```
GET /api/playback/session/{sessionID}

Response:
{
  "session_id": "uuid",
  "user_id": "user123",
  "profile_id": "profile456",
  "content_id": "content789",
  "device_id": "device101",
  "status": "active",
  "started_at": "2026-02-18T08:54:22Z",
  "last_heartbeat": "2026-02-18T08:54:22Z",
  "position": 0
}
```

#### Send Heartbeat
```
POST /api/playback/session/{sessionID}/heartbeat
Content-Type: application/json

{
  "position": 120  // playback position in seconds
}

Response:
{
  "status": "ok"
}
```

#### Pause Session
```
POST /api/playback/session/{sessionID}/pause

Response:
{
  "status": "paused"
}
```

#### Resume Session
```
POST /api/playback/session/{sessionID}/resume

Response:
{
  "status": "active"
}
```

#### Stop Session
```
POST /api/playback/session/{sessionID}/stop

Response:
{
  "status": "stopped"
}
```

#### Terminate Session
```
DELETE /api/playback/session/{sessionID}/terminate

Response:
{
  "status": "terminated"
}
```

### Token Management

#### Validate Token
```
POST /api/playback/token/validate
Content-Type: application/json

{
  "token": "jwt-token-here"
}

Response:
{
  "valid": true,
  "claims": {
    "session_id": "uuid",
    "user_id": "user123",
    "profile_id": "profile456",
    "content_id": "content789",
    "device_id": "device101"
  }
}
```

#### Refresh Token
```
POST /api/playback/token/refresh
Content-Type: application/json

{
  "token": "old-jwt-token"
}

Response:
{
  "token": "new-jwt-token"
}
```

### Concurrency Management

#### Get Concurrency Info
```
GET /api/concurrency/{userID}

Response:
{
  "user_id": "user123",
  "active_stream_count": 2,
  "max_concurrent_streams": 4,
  "active_sessions": ["session1", "session2"]
}
```

## Event Emission

The service emits the following events to the Engagement Service:
- `playback.started` - When a playback session is created
- `playback.paused` - When playback is paused
- `playback.resumed` - When playback is resumed
- `playback.stopped` - When playback is stopped
- `playback.completed` - When a session is terminated
- `playback.heartbeat` - On each heartbeat update

Events are currently logged. In production, these should be published to a message queue (Kafka, RabbitMQ, etc.) for the Engagement Service to consume.

## Design Patterns

### Stateless Design
- No local state stored in the service
- All session state stored in Redis
- Enables horizontal scaling

### Concurrency Control
- Distributed locking using Redis sets
- Atomic operations for slot acquisition
- Prevents race conditions

### Security
- JWT-based authentication tokens
- Short-lived tokens (configurable expiration)
- Token refresh mechanism

### Resilience
- Graceful degradation if external services are unavailable
- Timeout on external service calls
- Redis connection retry logic

## Dependencies

- `github.com/go-redis/redis/v8` - Redis client
- `github.com/gorilla/mux` - HTTP router
- `github.com/golang-jwt/jwt/v5` - JWT token generation
- `github.com/google/uuid` - UUID generation

## Future Enhancements

1. **Database Integration**: Implement PostgreSQL for durable session logs
2. **Event Bus Integration**: Connect to Kafka/RabbitMQ for event streaming
3. **Metrics & Monitoring**: Add Prometheus metrics
4. **Circuit Breakers**: Add resilience patterns for external service calls
5. **Rate Limiting**: Add per-user rate limiting
6. **Caching**: Add response caching for frequently accessed data
7. **Authentication**: Add authentication middleware for API endpoints

## License

This is part of the klones project.
