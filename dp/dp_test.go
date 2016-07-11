package dp

import (
	"reflect"
	"testing"
)

func TestCode338(t *testing.T) {

	kv := map[int][]int{
		0: []int{0},
		1: []int{0, 1},
		5: []int{0, 1, 1, 2, 1, 2},
	}

	for k, v := range kv {
		res := countBits(k)
		if !reflect.DeepEqual(res, v) {
			t.Errorf("Error @ (%d, %s) with output is %s", k, v, res)
		}
	}
	t.Log("Success")
}
