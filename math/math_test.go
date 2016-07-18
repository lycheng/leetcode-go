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

func TestCode231(t *testing.T) {
	if isPowerOfTwo(0) {
		t.Error("error")
	}
	if !isPowerOfTwo(1) {
		t.Error("error")
	}
	if !isPowerOfTwo(4) {
		t.Error("error")
	}
	if isPowerOfTwo(6) {
		t.Error("error")
	}
	t.Log("Success")
}

func TestCode342(t *testing.T) {
	if isPowerOfFour(0) {
		t.Error("error")
	}
	if !isPowerOfFour(1) {
		t.Error("error")
	}
	if !isPowerOfFour(4) {
		t.Error("error")
	}
	if isPowerOfFour(8) {
		t.Error("error")
	}
	t.Log("Success")
}

func TestCode263(t *testing.T) {
	if isUgly(0) {
		t.Error("error")
	}
	if !isUgly(1) {
		t.Error("error")
	}
	if isUgly(14) {
		t.Error("error")
	}
	t.Log("Success")
}
