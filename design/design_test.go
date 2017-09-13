package design

import (
	"testing"
)

type Code460Action struct {
	OP   string
	Args [2]int
}

func TestCode460(t *testing.T) {
	var cases []Code460Action

	cases = []Code460Action{
		Code460Action{"put", [2]int{2, 1}},
		Code460Action{"get", [2]int{2, 1}},
		Code460Action{"put", [2]int{3, 2}},
		Code460Action{"get", [2]int{2, -1}},
		Code460Action{"get", [2]int{3, 2}},
	}
	runCases(t, 1, cases)

	cases = []Code460Action{
		Code460Action{"put", [2]int{1, 1}},
		Code460Action{"put", [2]int{2, 2}},
		Code460Action{"get", [2]int{1, 1}},
		Code460Action{"put", [2]int{3, 3}},
		Code460Action{"get", [2]int{2, -1}},
		Code460Action{"get", [2]int{3, 3}},
		Code460Action{"put", [2]int{4, 4}},
		Code460Action{"get", [2]int{1, -1}},
		Code460Action{"get", [2]int{3, 3}},
		Code460Action{"get", [2]int{4, 4}},
	}
	runCases(t, 2, cases)

	cases = []Code460Action{
		Code460Action{"put", [2]int{3, 1}},
		Code460Action{"put", [2]int{2, 1}},
		Code460Action{"put", [2]int{2, 2}},
		Code460Action{"put", [2]int{4, 4}},
		Code460Action{"get", [2]int{2, 2}},
	}
	runCases(t, 2, cases)

	cases = []Code460Action{
		Code460Action{"get", [2]int{2, -1}},
		Code460Action{"put", [2]int{2, 6}},
		Code460Action{"get", [2]int{1, -1}},
		Code460Action{"put", [2]int{1, 5}},
		Code460Action{"put", [2]int{1, 2}},
		Code460Action{"get", [2]int{1, 2}},
		Code460Action{"get", [2]int{2, 6}},
	}
	runCases(t, 2, cases)
}

func runCases(t *testing.T, size int, cases []Code460Action) {
	cache := Constructor(size)
	for _, item := range cases {
		if item.OP == "put" {
			t.Logf("put (%d %d)", item.Args[0], item.Args[1])
			cache.Put(item.Args[0], item.Args[1])
		} else {
			t.Logf("get (%d %d)", item.Args[0], item.Args[1])
			r := cache.Get(item.Args[0])
			if r != item.Args[1] {
				t.Fatal("error @", item.Args, "result is", r)
			}
		}
	}
}
