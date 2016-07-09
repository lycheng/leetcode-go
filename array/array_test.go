package array

import (
	"reflect"
	"testing"
)

func TestCode238(t *testing.T) {

	src := []int{1, 2, 3, 4}
	dst := []int{24, 12, 8, 6}

	result := productExceptSelf(src)

	if !reflect.DeepEqual(dst, result) {
		t.Error("Error")
	}
	t.Log("Success")
}
