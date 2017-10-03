package math

import (
	"reflect"
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

func TestCode50(t *testing.T) {
	if myPow(2, 2) != 4.0 {
		t.Error("error")
	}

	if myPow(2, -2) != 0.25 {
		t.Error("error")
	}
	t.Log("Success")
}

func TestCode645(t *testing.T) {
	src := []int{1, 2, 2, 4}
	dst := []int{2, 3}

	ret := findErrorNums(src)
	if !reflect.DeepEqual(ret, dst) {
		t.Fatal("error")
	}

	src = []int{3, 2, 3, 4, 6, 5}
	dst = []int{3, 1}
	ret = findErrorNums(src)
	if !reflect.DeepEqual(ret, dst) {
		t.Fatal("error with not sorted array")
	}
	t.Log("success")
}
