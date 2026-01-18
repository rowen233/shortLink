package service

import (
	"testing"
)

func TestGenerateShortCode(t *testing.T) {
	service := &ShortLinkService{
		baseURL: "http://localhost:8080",
	}

	tests := []struct {
		name string
		url  string
	}{
		{
			name: "GitHub URL",
			url:  "https://github.com",
		},
		{
			name: "Google URL",
			url:  "https://www.google.com",
		},
		{
			name: "Long URL",
			url:  "https://www.example.com/very/long/path/to/resource?param1=value1&param2=value2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := service.generateShortCode(tt.url)

			if code == "" {
				t.Error("Generated short code should not be empty")
			}

			if len(code) > 8 {
				t.Errorf("Short code length should be <= 8, got %d", len(code))
			}

			t.Logf("URL: %s -> Short code: %s", tt.url, code)
		})
	}
}

func TestGenerateShortCodeConsistency(t *testing.T) {
	service := &ShortLinkService{
		baseURL: "http://localhost:8080",
	}

	url := "https://github.com"
	code1 := service.generateShortCode(url)
	code2 := service.generateShortCode(url)

	if code1 != code2 {
		t.Errorf("Same URL should generate same short code, got %s and %s", code1, code2)
	}
}
