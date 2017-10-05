package fun

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCode412(t *testing.T) {
	dst := []string{
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz",
	}
	if ret := fizzBuzz(15); !reflect.DeepEqual(ret, dst) {
		fmt.Println(ret)
		t.Fatal("error")
	}
	t.Log("success")
}

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
