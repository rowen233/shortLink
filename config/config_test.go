package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg := LoadConfig()

	if cfg.ServerPort == "" {
		t.Error("ServerPort should not be empty")
	}

	if cfg.RedisAddr == "" {
		t.Error("RedisAddr should not be empty")
	}

	if cfg.BaseURL == "" {
		t.Error("BaseURL should not be empty")
	}
}

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		setValue     string
		defaultValue string
		expected     string
	}{
		{
			name:         "Environment variable exists",
			key:          "TEST_VAR_1",
			setValue:     "test_value",
			defaultValue: "default",
			expected:     "test_value",
		},
		{
			name:         "Environment variable does not exist",
			key:          "TEST_VAR_2",
			setValue:     "",
			defaultValue: "default",
			expected:     "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setValue != "" {
				os.Setenv(tt.key, tt.setValue)
				defer os.Unsetenv(tt.key)
			}

			result := getEnv(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
