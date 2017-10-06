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

func TestCode463(t *testing.T) {
	var grid [][]int = [][]int{
		[]int{0, 1, 0, 0},
		[]int{1, 1, 1, 0},
		[]int{0, 1, 0, 0},
		[]int{1, 1, 0, 0},
	}

	if ret := islandPerimeter(grid); ret != 16 {
		t.Fatal("get", ret, "and expect", 16)
	}
	t.Log("success")
}
