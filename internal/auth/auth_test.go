package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectError bool
	}{
		{
			name:        "Valid",
			headers:     http.Header{"Authorization": []string{"ApiKey 1234"}}, // Fixed case
			expectError: false,
		},
		{
			name:        "Invalid",
			headers:     http.Header{"Authorization": []string{"Apikeyeiei 1234"}},
			expectError: true,
		},
		{
			name:        "No header",
			headers:     http.Header{},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := GetAPIKey(tc.headers)

			if tc.expectError && err == nil {
				t.Error("expected error but got none")
			}
			if !tc.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}
		})
	}
}
