package model

import (
	"encoding/json"
	"testing"
)

func TestTimeframe_JSON(t *testing.T) {
	tests := []struct {
		name      string
		timeframe Timeframe
		expected  string
	}{
		{
			name:      "valid timeframe",
			timeframe: Timeframe{N: 5, Unit: "minutes"},
			expected:  `{"n":5,"unit":"minutes"}`,
		},
		{
			name:      "zero value",
			timeframe: Timeframe{N: 0, Unit: ""},
			expected:  `{"n":0,"unit":""}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.timeframe)
			if err != nil {
				t.Errorf("Failed to marshal timeframe: %v", err)
			}
			if string(data) != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, string(data))
			}
		})
	}
}

func TestTimeframe_Unmarshal(t *testing.T) {
	jsonData := `{"n":10,"unit":"seconds"}`
	var tf Timeframe

	err := json.Unmarshal([]byte(jsonData), &tf)
	if err != nil {
		t.Errorf("Failed to unmarshal timeframe: %v", err)
	}

	if tf.N != 10 {
		t.Errorf("Expected N=10, got %d", tf.N)
	}
	if tf.Unit != "seconds" {
		t.Errorf("Expected Unit=seconds, got %s", tf.Unit)
	}
}
