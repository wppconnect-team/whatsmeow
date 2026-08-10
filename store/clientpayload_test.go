package store

import "testing"

func TestParseVersionBounds(t *testing.T) {
	parsed, err := ParseVersion("2.3000.123")
	if err != nil {
		t.Fatalf("expected valid version: %v", err)
	}
	if parsed != (WAVersionContainer{2, 3000, 123}) {
		t.Fatalf("unexpected parsed version: %v", parsed)
	}

	for _, invalid := range []string{"-1.2.3", "4294967296.2.3"} {
		if _, err = ParseVersion(invalid); err == nil {
			t.Fatalf("expected %q to be rejected", invalid)
		}
	}
}
