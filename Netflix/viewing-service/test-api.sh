#!/bin/bash

# Viewing Service API Test Script
# This script demonstrates the basic API usage of the Viewing Service

BASE_URL="http://localhost:8080"

echo "========================================="
echo "Viewing Service API Test"
echo "========================================="
echo ""

# 1. Health Check
echo "1. Health Check"
echo "   GET $BASE_URL/health"
curl -s -X GET "$BASE_URL/health" | python3 -m json.tool
echo ""
echo ""

# 2. Create a Playback Session
echo "2. Create Playback Session"
echo "   POST $BASE_URL/api/playback/session"
SESSION_RESPONSE=$(curl -s -X POST "$BASE_URL/api/playback/session" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user123",
    "profile_id": "profile456",
    "content_id": "content789",
    "device_id": "device101"
  }')

echo "$SESSION_RESPONSE" | python3 -m json.tool

# Extract session ID and token
SESSION_ID=$(echo "$SESSION_RESPONSE" | python3 -c "import sys, json; print(json.load(sys.stdin).get('session', {}).get('session_id', ''))" 2>/dev/null || echo "")
PLAYBACK_TOKEN=$(echo "$SESSION_RESPONSE" | python3 -c "import sys, json; print(json.load(sys.stdin).get('playback_token', ''))" 2>/dev/null || echo "")

if [ -z "$SESSION_ID" ]; then
    echo ""
    echo "⚠️  Warning: Session creation may have failed (likely because Redis is not running)"
    echo "   The service needs Redis for session storage."
    echo "   To run Redis: docker run -d -p 6379:6379 redis:alpine"
    echo ""
    echo "Skipping remaining tests that require a valid session..."
    exit 0
fi

echo ""
echo "Session ID: $SESSION_ID"
echo "Playback Token: $PLAYBACK_TOKEN"
echo ""
echo ""

# 3. Get Session Details
echo "3. Get Session Details"
echo "   GET $BASE_URL/api/playback/session/$SESSION_ID"
curl -s -X GET "$BASE_URL/api/playback/session/$SESSION_ID" | python3 -m json.tool
echo ""
echo ""

# 4. Send Heartbeat
echo "4. Send Heartbeat (position: 120 seconds)"
echo "   POST $BASE_URL/api/playback/session/$SESSION_ID/heartbeat"
curl -s -X POST "$BASE_URL/api/playback/session/$SESSION_ID/heartbeat" \
  -H "Content-Type: application/json" \
  -d '{
    "position": 120
  }' | python3 -m json.tool
echo ""
echo ""

# 5. Pause Session
echo "5. Pause Session"
echo "   POST $BASE_URL/api/playback/session/$SESSION_ID/pause"
curl -s -X POST "$BASE_URL/api/playback/session/$SESSION_ID/pause" | python3 -m json.tool
echo ""
echo ""

# 6. Resume Session
echo "6. Resume Session"
echo "   POST $BASE_URL/api/playback/session/$SESSION_ID/resume"
curl -s -X POST "$BASE_URL/api/playback/session/$SESSION_ID/resume" | python3 -m json.tool
echo ""
echo ""

# 7. Validate Token
echo "7. Validate Playback Token"
echo "   POST $BASE_URL/api/playback/token/validate"
curl -s -X POST "$BASE_URL/api/playback/token/validate" \
  -H "Content-Type: application/json" \
  -d "{
    \"token\": \"$PLAYBACK_TOKEN\"
  }" | python3 -m json.tool
echo ""
echo ""

# 8. Get Concurrency Info
echo "8. Get Concurrency Info"
echo "   GET $BASE_URL/api/concurrency/user123"
curl -s -X GET "$BASE_URL/api/concurrency/user123" | python3 -m json.tool
echo ""
echo ""

# 9. Stop Session
echo "9. Stop Session"
echo "   POST $BASE_URL/api/playback/session/$SESSION_ID/stop"
curl -s -X POST "$BASE_URL/api/playback/session/$SESSION_ID/stop" | python3 -m json.tool
echo ""
echo ""

# 10. Terminate Session
echo "10. Terminate Session"
echo "    DELETE $BASE_URL/api/playback/session/$SESSION_ID/terminate"
curl -s -X DELETE "$BASE_URL/api/playback/session/$SESSION_ID/terminate" | python3 -m json.tool
echo ""
echo ""

echo "========================================="
echo "Test Complete!"
echo "========================================="
