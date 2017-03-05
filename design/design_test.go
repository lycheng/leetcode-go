package design

import (
	"testing"
)

type OP struct {
	OP   string
	Args [2]int
}

func TestCode460(t *testing.T) {
	cache := Constructor(2)
	items := []OP{
		OP{"put", [2]int{1, 1}},
		OP{"put", [2]int{2, 2}},
		OP{"get", [2]int{1, 1}},
		OP{"put", [2]int{3, 3}},
		OP{"get", [2]int{2, -1}},
		OP{"get", [2]int{3, 3}},
		OP{"put", [2]int{4, 4}},
		OP{"get", [2]int{1, -1}},
		OP{"get", [2]int{3, 3}},
		OP{"get", [2]int{4, 4}},
	}
	for _, item := range items {
		if item.OP == "put" {
			cache.Put(item.Args[0], item.Args[1])
		} else {
			r := cache.Get(item.Args[0])
			if r != item.Args[1] {
				t.Fatal("error @", item.Args, "result is", r)
			}
		}

	}
}
