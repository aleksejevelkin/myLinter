package checkers

import "testing"

func TestCheckNoSensitiveDataWithKeywords(t *testing.T) {
	t.Run("nil keywords -> default list", func(t *testing.T) {
		if err := CheckNoSensitiveDataWithKeywords("password: 123", nil); err == nil {
			t.Fatalf("expected error for default keyword list")
		}
	})

	t.Run("custom list overrides default", func(t *testing.T) {
		// В дефолтном списке есть password, но мы передаём только "pin".
		if err := CheckNoSensitiveDataWithKeywords("password: 123", []string{"pin"}); err != nil {
			t.Fatalf("did not expect error, got %v", err)
		}

		if err := CheckNoSensitiveDataWithKeywords("pin: 123", []string{"pin"}); err == nil {
			t.Fatalf("expected error for custom keyword list")
		}
	})
}
