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
