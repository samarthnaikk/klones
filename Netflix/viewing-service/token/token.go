package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Generator handles JWT token generation
type Generator struct {
	secret     []byte
	expiration time.Duration
}

// PlaybackClaims represents the JWT claims for playback tokens
type PlaybackClaims struct {
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id"`
	ProfileID string `json:"profile_id"`
	ContentID string `json:"content_id"`
	DeviceID  string `json:"device_id"`
	jwt.RegisteredClaims
}

// NewGenerator creates a new token generator
func NewGenerator(secret string, expiration time.Duration) *Generator {
	return &Generator{
		secret:     []byte(secret),
		expiration: expiration,
	}
}

// GeneratePlaybackToken generates a JWT token for playback
func (g *Generator) GeneratePlaybackToken(sessionID, userID, profileID, contentID, deviceID string) (string, error) {
	now := time.Now()
	claims := PlaybackClaims{
		SessionID: sessionID,
		UserID:    userID,
		ProfileID: profileID,
		ContentID: contentID,
		DeviceID:  deviceID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(g.expiration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "viewing-service",
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(g.secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken validates a playback token and returns the claims
func (g *Generator) ValidateToken(tokenString string) (*PlaybackClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &PlaybackClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return g.secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*PlaybackClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// RefreshToken generates a new token from an existing valid token
func (g *Generator) RefreshToken(oldTokenString string) (string, error) {
	claims, err := g.ValidateToken(oldTokenString)
	if err != nil {
		return "", fmt.Errorf("cannot refresh invalid token: %w", err)
	}

	// Generate new token with same claims but updated expiration
	return g.GeneratePlaybackToken(claims.SessionID, claims.UserID, claims.ProfileID, claims.ContentID, claims.DeviceID)
}
