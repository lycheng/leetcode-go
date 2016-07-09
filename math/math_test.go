package math

import (
	"testing"
)

func TestCode258(t *testing.T) {

	if addDigits(9) != 9 {
		t.Error("error")
	}

	if addDigits(25) != 7 {
		t.Error("error")
	}

	if addDigits(0) != 0 {
		t.Error("error")
	}
	t.Log("Success")
}
