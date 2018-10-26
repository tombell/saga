package decks_test

import (
	"testing"

	"github.com/tombell/saga/decks"
)

func TestStatusString(t *testing.T) {
	tests := []struct {
		name     string
		input    decks.Status
		expected string
	}{
		{"empty", decks.Empty, "EMPTY"},
		{"new", decks.New, "NEW"},
		{"playing", decks.Playing, "PLAYING"},
		{"played", decks.Played, "PLAYED"},
		{"skipped", decks.Skipped, "SKIPPED"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.String()
			expected := tc.expected

			if actual != expected {
				t.Fatalf("expected string to be %s, got %s", expected, actual)
			}
		})
	}
}
