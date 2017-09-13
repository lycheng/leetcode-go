// package array test codes
package array

import (
	"math/rand"
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

func BenchmarkCode238(t *testing.B) {
	srcLen := 1000

	src := make([]int, srcLen)
	for i := 0; i < srcLen; i++ {
		src[i] = rand.Int()
	}

	for i := 0; i < t.N; i++ {
		productExceptSelf(src)
	}
}

func TestCode189(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6, 7}
	dst := []int{5, 6, 7, 1, 2, 3, 4}

	rotate(src, 3)

	if !reflect.DeepEqual(src, dst) {
		t.Error("Error")
	}

	src = []int{2147483647, -2147483648, 33, 219, 0}
	dst = []int{-2147483648, 33, 219, 0, 2147483647}
	rotate(src, 4)
	if !reflect.DeepEqual(src, dst) {
		t.Error("Error")
	}

	t.Log("Success")
}
