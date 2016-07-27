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

func TestCode20(t *testing.T) {
	validStrings := []string{
		"()",
		"()[]{}",
	}

	for _, s := range validStrings {
		if !isValid(s) {
			t.Error("error")
		}
	}

	notValidStrings := []string{
		"}",
		"{",
		"(]",
		"([)]",
	}
	for _, s := range notValidStrings {
		if isValid(s) {
			t.Error("error")
		}
	}
	t.Log("success")
}
