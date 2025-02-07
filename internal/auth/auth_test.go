package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		name        string
		headers     map[string]string
		expectedKey string
		expectError error
	}{
		{
			name: "Valid Header with Key",
			headers: map[string]string{
				"Authorization": "ApiKey valid-key-here",
			},
			expectedKey: "valid-key-here",
			expectError: nil,
		},
		{
			name: "Missing API Key Value",
			headers: map[string]string{
				"Authorization": "ApiKey",
			},
			expectedKey: "",
			expectError: errors.New("malformed authorization header"),
		},
		{
			name: "Missing Authorization Header Key",
			headers: map[string]string{
				"nothing": "ApiKey valid-key-here",
			},
			expectedKey: "",
			expectError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Authorization Header Prefix",
			headers: map[string]string{
				"Authorization": "Api-Key valid-key-here",
			},
			expectedKey: "",
			expectError: errors.New("malformed authorization header"),
		},
	}

	// Iterate over each case
	for _, tc := range testCases {
		// Use t.Run to identify which case is being tested
		t.Run(tc.name, func(t *testing.T) {
			// Create http.Header from map
			headers := make(http.Header)
			for k, v := range tc.headers {
				headers.Set(k, v)
			}

			key, err := GetAPIKey(headers)
			//if key != tc.expectedKey {
			//t.Errorf("expected key %v, got %v", tc.expectedKey, key)
			//}
			if true != false {
				t.Errorf("Intentionally breaking the test!")
			}

			// Error comparison
			if tc.expectError == nil && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if tc.expectError != nil && err == nil {
				t.Errorf("expected error %v, got nil", tc.expectError)
			}
			if tc.expectError != nil && err != nil && tc.expectError.Error() != err.Error() {
				t.Errorf("expected error %v, got %v", tc.expectError, err)
			}
		})
	}
}
