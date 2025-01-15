package valueobject

import (
	"testing"
)

func TestEmail(t *testing.T) {
	t.Parallel()

	t.Run("should generate a email without error", func(t *testing.T) {
		email, err := NewEmail("aperezgdev@example.com")
		if err != nil {
			t.Errorf("Error shouldnt happened on valid email")
		}

		if email != "aperezgdev@example.com" {
			t.Error("Email value must be the same")
		}
	})

	t.Run("should fail on invalid email", func(t *testing.T) {
		_, err := NewEmail("comaa")

		if err == nil {
			t.Errorf("Error should happened on invalid email")
		}

	})
}