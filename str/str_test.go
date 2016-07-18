package str

import (
	"testing"
)

func TestCode344(t *testing.T) {

	src := "hello"
	dst := "olleh"

	if reverseString(src) != dst {
		t.Error("error")
	}
	t.Log("success")
}

func TestCode345(t *testing.T) {

	src := "hello"
	dst := "holle"
	if reverseVowels(src) != dst {
		t.Error("error")
	}

	src = "leetcode"
	dst = "leotcede"
	if reverseVowels(src) != dst {
		t.Error("error")
	}
	t.Log("success")
}
