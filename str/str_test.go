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

func TestCode28(t *testing.T) {
	src := "hello"
	target := "ello"
	result := 1

	if strStr(src, target) != result {
		t.Error("error")
	}

	src = "hello,world"
	target = "abc"
	result = -1

	if strStr(src, target) != result {
		t.Error("error")
	}
	t.Log("success")

	src = "hello,world"
	target = ""
	result = 0

	if strStr(src, target) != result {
		t.Error("error")
	}

	src = "hello"
	target = "hello"
	result = 0

	if strStr(src, target) != result {
		t.Error("error")
	}
	t.Log("success")
}
