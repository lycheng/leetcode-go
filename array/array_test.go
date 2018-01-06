// package array test codes
package array

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestCode167(t *testing.T) {
	src := []int{2, 7, 11, 15}
	result := twoSum(src, 9)
	if result[0] != 1 && result[1] != 2 {
		t.Fatal("error")
	}
	t.Log("success")
}

func TestCode189(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6, 7}
	dst := []int{5, 6, 7, 1, 2, 3, 4}

	rotate(src, 3)

	if !reflect.DeepEqual(src, dst) {
		t.Error("error")
	}

	src = []int{2147483647, -2147483648, 33, 219, 0}
	dst = []int{-2147483648, 33, 219, 0, 2147483647}
	rotate(src, 4)
	if !reflect.DeepEqual(src, dst) {
		t.Error("error")
	}

	t.Log("success")
}

func TestCode238(t *testing.T) {

	src := []int{1, 2, 3, 4}
	dst := []int{24, 12, 8, 6}

	result := productExceptSelf(src)

	if !reflect.DeepEqual(dst, result) {
		t.Error("error")
	}
	t.Log("success")
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

func TestCode485(t *testing.T) {
	src := []int{1, 1, 0, 1, 1, 1}
	if l := findMaxConsecutiveOnes(src); l != 3 {
		t.Error("got", l, "and expect 3")
	}
	t.Log("success")
}

func TestCode561(t *testing.T) {
	src := []int{1, 4, 3, 2}
	if n := arrayPairSum(src); n != 4 {
		t.Error("error")
	}
	t.Log("success")
}

func TestCode674(t *testing.T) {
	src := []int{1, 3, 5, 4, 7}
	if l := findLengthOfLCIS(src); l != 3 {
		t.Error("error with", l, "and expect 4")
	}

	src = []int{1, 3, 5, 7}
	if l := findLengthOfLCIS(src); l != 4 {
		t.Error("error with", l, "and expect 4")
	}
	t.Log("success")
}

func TestCode744(t *testing.T) {
	src := []byte{
		byte('c'), byte('f'), byte('j'),
	}
	target := byte('j')
	if r := nextGreatestLetter(src, target); r != byte('c') {
		t.Fatal("got", string(r), "and expect c")
	}

	src = []byte{
		byte('c'), byte('f'), byte('j'),
	}
	target = byte('a')
	if r := nextGreatestLetter(src, target); r != byte('c') {
		t.Fatal("got", string(r), "and expect c")
	}

	src = []byte{
		byte('c'), byte('f'), byte('j'),
	}
	target = byte('c')
	if r := nextGreatestLetter(src, target); r != byte('f') {
		t.Fatal("got", string(r), "and expect f")
	}

	t.Log("success")
}
