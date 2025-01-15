package valueobject

import (
	"strings"
	"testing"
)

func TestTitle(t *testing.T) {
	t.Parallel()

	t.Run("should generate title without error", func(t *testing.T) {
		title, err := NewTitle("titulo valido")
		if err != nil {
			t.Errorf("Error shouldnt happened on valid title value")
		}
		if title != "titulo valido" {
			t.Errorf("Title value is not correct")
		}
	})

	t.Run("should return error on empty title value", func(t *testing.T) {
		_, err := NewTitle("")
		if err == nil {
			t.Errorf("Error should happened on empty value")
		}
	})

	t.Run("should return error on too long title value", func(t *testing.T) {
		_, err := NewTitle(strings.Repeat("a", 41))
		if err == nil {
			t.Errorf("Error should happened too long title value")
		}
	})
}