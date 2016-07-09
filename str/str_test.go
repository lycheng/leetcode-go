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
