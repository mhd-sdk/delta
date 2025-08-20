package model

import (
	"encoding/json"
	"testing"
)

func TestUpdateStateRequest_JSON(t *testing.T) {
	tests := []struct {
		name     string
		request  UpdateStateRequest
		expected string
	}{

		{
			name:     "valid request",
			request:  UpdateStateRequest{ID: "123", Status: "active"},
			expected: `{"id":"123","status":"active"}`,
		},
		{
			name:     "empty values",
			request:  UpdateStateRequest{ID: "", Status: ""},
			expected: `{"id":"","status":""}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.request)
			if err != nil {
				t.Errorf("Failed to marshal request: %v", err)
			}
			if string(data) != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, string(data))
			}
		})
	}
}
