package auth

import (
	"testing"
	"time"
)

func TestIssueParse(t *testing.T) {
	m := NewTokenManager("test-secret-at-least-16", 24*time.Hour)
	tok, err := m.Issue(42)
	if err != nil {
		t.Fatal(err)
	}
	id, err := m.Parse(tok)
	if err != nil {
		t.Fatal(err)
	}
	if id != 42 {
		t.Fatalf("got %d", id)
	}
}
