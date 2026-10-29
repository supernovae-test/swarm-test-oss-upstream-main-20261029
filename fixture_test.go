package fixture

import "testing"

func TestValue(t *testing.T) {
	if Value() != "fixture" {
		t.Fatal("unexpected fixture value")
	}
}
