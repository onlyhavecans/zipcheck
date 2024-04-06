package ziptools_test

import (
	"testing"

	"github.com/onlyhavecans/zipcheck/internal/ziptools"
)

func TestIsValidZip(t *testing.T) {
	// test if isValidZip returns true for a valid zip file
	// and false for an invalid zip file
	if !ziptools.IsValidZip("../../testdata/good.zip") {
		t.Errorf("Expected true for valid.zip")
	}
	if ziptools.IsValidZip("../../testdata/bad01.zip") {
		t.Errorf("Expected false for invalid.zip")
	}
	if ziptools.IsValidZip("../../testdata/incomplete.zip") {
		t.Errorf("Expected false for incomplete.zip")
	}
}
