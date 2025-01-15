package valueobject

import "testing"

func TestId(t *testing.T) {
	t.Parallel()

	t.Run("should parse valid id", func(t *testing.T) {
		_, err := ValidateId("01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err != nil {
			t.Errorf("Error shouldnt happened on valid id")
		}
	})

	t.Run("should fail on invalid id", func(t *testing.T) {
		_, err := ValidateId("a")
		if err == nil {
			t.Errorf("Error should happened on invalid id")
		}
	})

	t.Run("should generate id without error", func(t *testing.T) {
		_, err := NewId()
		if err != nil {
			t.Errorf("Error shouldnt happened generating id")
		}
	})
}