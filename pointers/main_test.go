package pointers_test

import (
	"cards/pointers"
	"testing"
)

func TestName(t *testing.T) {
	if !pointers.Example() {
		t.Errorf("Expected true, but got false")
	}
}
