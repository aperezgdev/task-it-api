package valueobject

import (
	"strings"
	"testing"
)

func TestDescription(t *testing.T) {
	t.Parallel()

	t.Run("should generate valid description on valid value", func(t *testing.T) {
		description, err := NewDescription("description")
		if err != nil {
			t.Errorf("Error shouldnt happened on valid description value")
		}

		if description != "description" {
			t.Errorf("Description value is not correct")
		}
	})
	
	t.Run("should fail on too long description", func(t *testing.T) {
		_, err := NewDescription(strings.Repeat("a", 241))
		if err == nil {
			t.Errorf("Error should happened on too long description")
		}
	})
}