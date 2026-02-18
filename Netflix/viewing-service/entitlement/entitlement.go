package entitlement

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Validator handles entitlement validation
type Validator struct {
	identityServiceURL string
	httpClient         *http.Client
}

// EntitlementResponse represents the response from Identity Service
type EntitlementResponse struct {
	Entitled bool   `json:"entitled"`
	Reason   string `json:"reason,omitempty"`
}

// NewValidator creates a new entitlement validator
func NewValidator(identityServiceURL string) *Validator {
	return &Validator{
		identityServiceURL: identityServiceURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// ValidateEntitlement checks if a user is entitled to view content
func (v *Validator) ValidateEntitlement(ctx context.Context, userID, profileID, contentID string) (bool, error) {
	// Construct the request to Identity Service
	url := fmt.Sprintf("%s/api/entitlement/validate?userID=%s&profileID=%s&contentID=%s",
		v.identityServiceURL, userID, profileID, contentID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := v.httpClient.Do(req)
	if err != nil {
		// If Identity Service is unavailable, we could implement fallback logic here
		// For now, we'll deny access
		return false, fmt.Errorf("failed to call Identity Service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return false, fmt.Errorf("Identity Service returned error: %s (status: %d)", body, resp.StatusCode)
	}

	var entitlementResp EntitlementResponse
	err = json.NewDecoder(resp.Body).Decode(&entitlementResp)
	if err != nil {
		return false, fmt.Errorf("failed to decode response: %w", err)
	}

	if !entitlementResp.Entitled {
		return false, fmt.Errorf("user not entitled: %s", entitlementResp.Reason)
	}

	return true, nil
}

// ValidateContentAvailability checks if content is available
func (v *Validator) ValidateContentAvailability(ctx context.Context, contentID, region string, contentServiceURL string) (bool, error) {
	// Construct request to Content Service
	url := fmt.Sprintf("%s/api/content/availability?contentID=%s&region=%s",
		contentServiceURL, contentID, region)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		// If Content Service is unavailable, we could implement fallback logic
		return false, fmt.Errorf("failed to call Content Service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("content not available (status: %d)", resp.StatusCode)
	}

	var availabilityResp struct {
		Available bool   `json:"available"`
		Reason    string `json:"reason,omitempty"`
	}
	
	err = json.NewDecoder(resp.Body).Decode(&availabilityResp)
	if err != nil {
		return false, fmt.Errorf("failed to decode response: %w", err)
	}

	return availabilityResp.Available, nil
}
