package model

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	if ErrNotFound == nil {
		t.Error("ErrNotFound should not be nil")
	}

	if ErrNotFound.Error() != "not found" {
		t.Errorf("Expected 'not found', got '%s'", ErrNotFound.Error())
	}
}

func TestErrorComparison(t *testing.T) {
	err := errors.New("not found")
	if err.Error() != ErrNotFound.Error() {
		t.Error("Error messages should match")
	}

	// Test that they are different error instances
	if err == ErrNotFound {
		t.Error("Should be different error instances")
	}
}
