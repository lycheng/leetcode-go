package fun

import "testing"

func TestCode657(t *testing.T) {
	moves := "ll"
	if judgeCircle(moves) {
		t.Fatal(moves)
	}
	moves = "UD"
	if !judgeCircle(moves) {
		t.Fatal(moves)
	}
	t.Log("Success")
}
