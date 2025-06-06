package fortune

import (
	"slices"
	"testing"
)

func TestGet(t *testing.T) {
	msg := Get()
	if i := slices.Index(fortunes, msg); i < 0 {
		t.Errorf("Expected one of %v, got %s", fortunes, msg)
	}
}

func TestEmpty(t *testing.T) {
	msg := Get()
	if msg == "" {
		t.Error("Expected a non-empty message, got an empty string")
	}
}
