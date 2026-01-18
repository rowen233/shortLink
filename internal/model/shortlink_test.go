package model

import (
	"testing"
	"time"
)

func TestShortLinkModel(t *testing.T) {
	shortLink := &ShortLink{
		ShortCode:   "abc123",
		OriginalURL: "https://example.com",
		CreatedAt:   time.Now(),
		VisitCount:  0,
	}

	if shortLink.ShortCode == "" {
		t.Error("ShortCode should not be empty")
	}

	if shortLink.OriginalURL == "" {
		t.Error("OriginalURL should not be empty")
	}

	if shortLink.VisitCount < 0 {
		t.Error("VisitCount should not be negative")
	}
}
