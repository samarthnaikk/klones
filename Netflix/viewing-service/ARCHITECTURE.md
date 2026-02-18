# Viewing Service Architecture

## System Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                         Client Applications                      │
│                    (Web, Mobile, Smart TV, etc.)                 │
└──────────────────────────┬──────────────────────────────────────┘
                           │
                           │ HTTP/REST
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Viewing Service API                         │
│                        (Port 8080)                               │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │ HTTP Router (Gorilla Mux)                                │  │
│  │  - Session endpoints                                      │  │
│  │  - Token endpoints                                        │  │
│  │  - Concurrency endpoints                                  │  │
│  │  - Health check                                           │  │
│  └──────────────────────────────────────────────────────────┘  │
└───────┬────────┬─────────┬─────────┬──────────┬────────────┬───┘
        │        │         │         │          │            │
        │        │         │         │          │            │
        ▼        ▼         ▼         ▼          ▼            ▼
     ┌────┐  ┌────┐   ┌────┐    ┌────┐    ┌────┐      ┌────────┐
     │Sess│  │Ent │   │Con │    │Tok │    │Lif │      │Event   │
     │ion │  │itle│   │curr│    │en  │    │ecyc│      │Emitter │
     │Mgr │  │ment│   │ency│    │Gen │    │le  │      │        │
     └─┬──┘  └──┬─┘   └──┬─┘    └────┘    └──┬─┘      └────┬───┘
       │        │        │                    │              │
       │        │        │                    │              │
       ▼        │        ▼                    ▼              ▼
    ┌──────────────────────┐              ┌────────────────────┐
    │       Redis          │              │   Event Stream     │
    │                      │              │   (Future: Kafka)  │
    │ - Session Storage    │              │                    │
    │ - Concurrency Locks  │              └────────────────────┘
    │ - Heartbeat Tracking │
    └──────────────────────┘
       ▲
       │
       │ External Service Calls (HTTP)
       │
       ├──────────────────────────────────┐
       │                                   │
       ▼                                   ▼
┌──────────────┐                   ┌──────────────┐
│   Identity   │                   │   Content    │
│   Service    │                   │   Service    │
│              │                   │              │
│ - Validate   │                   │ - Content    │
│   Entitlement│                   │   Availability│
└──────────────┘                   └──────────────┘
```

## Module Interaction Flow

### 1. Session Creation Flow

```
Client Request
    │
    ▼
┌────────────────────────────────────────────────┐
│ ViewingService.handleCreateSession()           │
└────────┬───────────────────────────────────────┘
         │
         ├─► 1. entitlement.ValidateEntitlement()
         │      └─► HTTP Call to Identity Service
         │
         ├─► 2. entitlement.ValidateContentAvailability()
         │      └─► HTTP Call to Content Service
         │
         ├─► 3. session.CreateSession()
         │      └─► Store in Redis
         │
         ├─► 4. concurrency.AcquireSlot()
         │      └─► Redis Set Operation
         │
         ├─► 5. token.GeneratePlaybackToken()
         │      └─► Create JWT
         │
         └─► 6. event.EmitPlaybackStarted()
                └─► Log/Queue Event
```

### 2. Heartbeat Update Flow

```
Client Heartbeat
    │
    ▼
┌────────────────────────────────────────────────┐
│ ViewingService.handleHeartbeat()               │
└────────┬───────────────────────────────────────┘
         │
         ├─► 1. session.GetSession()
         │      └─► Read from Redis
         │
         ├─► 2. lifecycle.HandleHeartbeat()
         │      └─► session.UpdateHeartbeat()
         │             └─► Update Redis
         │
         └─► 3. event.EmitHeartbeat()
                └─► Log/Queue Event
```

### 3. Session Termination Flow

```
Client Terminate
    │
    ▼
┌────────────────────────────────────────────────┐
│ ViewingService.handleTerminate()               │
└────────┬───────────────────────────────────────┘
         │
         ├─► 1. session.GetSession()
         │      └─► Read from Redis
         │
         ├─► 2. concurrency.ReleaseSlot()
         │      └─► Redis Set Remove
         │
         ├─► 3. session.TerminateSession()
         │      └─► Delete from Redis
         │
         └─► 4. event.EmitPlaybackCompleted()
                └─► Log/Queue Event
```

## Data Flow

### Session State Storage (Redis)

```
Key Pattern: session:{session_id}
Value: JSON {
    "session_id": "uuid",
    "user_id": "user123",
    "profile_id": "profile456",
    "content_id": "content789",
    "device_id": "device101",
    "status": "active|paused|stopped",
    "started_at": "2026-02-18T...",
    "last_heartbeat": "2026-02-18T...",
    "position": 120
}
TTL: 24h (configurable)
```

### Concurrency Tracking (Redis)

```
Key Pattern: concurrency:{user_id}
Type: SET
Members: [session_id1, session_id2, ...]
Operations:
  - SADD: Add session (acquire slot)
  - SREM: Remove session (release slot)
  - SCARD: Count active sessions
TTL: 24h (safety mechanism)
```

### User Session Index (Redis)

```
Key Pattern: user:{user_id}:sessions
Type: SET
Members: [session_id1, session_id2, ...]
Purpose: Quick lookup of user's active sessions
```

## JWT Token Structure

```json
{
  "header": {
    "alg": "HS256",
    "typ": "JWT"
  },
  "payload": {
    "session_id": "uuid",
    "user_id": "user123",
    "profile_id": "profile456",
    "content_id": "content789",
    "device_id": "device101",
    "iss": "viewing-service",
    "sub": "user123",
    "iat": 1708247422,
    "exp": 1708251022,
    "nbf": 1708247422
  },
  "signature": "HMACSHA256(...)"
}
```

## Event Schema

```json
{
  "event_type": "playback.started|paused|resumed|stopped|completed|heartbeat",
  "session_id": "uuid",
  "user_id": "user123",
  "profile_id": "profile456",
  "content_id": "content789",
  "device_id": "device101",
  "timestamp": "2026-02-18T08:54:22Z",
  "metadata": {
    "position": 120,
    "quality": "1080p",
    "network": "wifi"
  }
}
```

## Scaling Architecture

### Horizontal Scaling

```
         Load Balancer
              │
     ┌────────┼────────┐
     ▼        ▼        ▼
  [VS-1]  [VS-2]  [VS-3]  ... (Viewing Service Instances)
     │        │        │
     └────────┼────────┘
              │
         Redis Cluster
      (Master + Replicas)
```

### High Availability Setup

```
Region A                    Region B
┌─────────────┐            ┌─────────────┐
│ VS Instances│◄──────────►│ VS Instances│
│             │   Sync     │             │
└──────┬──────┘            └──────┬──────┘
       │                          │
       ▼                          ▼
┌─────────────┐            ┌─────────────┐
│Redis Master │◄──────────►│Redis Replica│
│             │  Replication│            │
└─────────────┘            └─────────────┘
```

## Security Layers

```
1. Client → HTTPS (TLS) → Load Balancer
2. Load Balancer → HTTP → Viewing Service
3. Viewing Service → JWT Validation → Request Processing
4. Viewing Service → HTTPS → Identity/Content Services
5. Viewing Service → AUTH (optional) → Redis
```

## Monitoring & Observability

```
Viewing Service
    │
    ├─► Logs → ELK Stack
    │
    ├─► Metrics → Prometheus
    │
    ├─► Traces → Jaeger (Future)
    │
    └─► Events → Kafka → Engagement Service
```

## Configuration Management

```
Environment Variables
    │
    ├─► Server Config (PORT)
    ├─► Redis Config (ADDR, PASSWORD, DB)
    ├─► PostgreSQL Config (optional)
    ├─► JWT Config (SECRET, EXPIRATION)
    ├─► Service URLs (IDENTITY_URL, CONTENT_URL)
    ├─► Limits (MAX_CONCURRENT_STREAMS)
    └─► Timeouts (SESSION_TIMEOUT, HEARTBEAT_INTERVAL)
```

## API Request Flow Example

### Complete Session Creation

```
1. POST /api/playback/session
   {
     "user_id": "user123",
     "profile_id": "profile456",
     "content_id": "content789",
     "device_id": "device101"
   }

2. Viewing Service validates input

3. Call Identity Service:
   GET http://identity-service/api/entitlement/validate
   ?userID=user123&profileID=profile456&contentID=content789

4. Call Content Service:
   GET http://content-service/api/content/availability
   ?contentID=content789&region=US

5. Create session in Redis:
   SET session:uuid {...} EX 86400

6. Acquire concurrency slot:
   SADD concurrency:user123 uuid

7. Generate JWT token

8. Emit event:
   LOG playback.started {...}

9. Return response:
   {
     "session": {...},
     "playback_token": "eyJ..."
   }
```

## Performance Characteristics

### Latency Targets
- Session creation: < 100ms (p95)
- Token generation: < 10ms (p95)
- Heartbeat update: < 50ms (p95)
- Concurrency check: < 20ms (p95)

### Throughput Targets
- Session creates: 10,000 RPS per instance
- Heartbeats: 50,000 RPS per instance
- Token validations: 100,000 RPS per instance

### Resource Usage (per instance)
- Memory: 50-100 MB
- CPU: 0.1-0.5 cores (idle-active)
- Network: 10-100 Mbps
- Redis connections: 10-50

This architecture ensures high availability, horizontal scalability, and low latency for Netflix-level streaming operations.
