package hash

import (
	"reflect"
	"testing"
)

func TestCode349(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	res := intersection(nums1, nums2)
	if len(res) != 1 || res[0] != 2 {
		t.Error("error")
	}
	t.Log("success")
}

func TestCode350(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	res := intersect(nums1, nums2)

	if !reflect.DeepEqual(nums2, res) {
		t.Error("error")
	}
	t.Log("success")
}

func TestCode202(t *testing.T) {

	if isHappy(0) {
		t.Error("error")
	}

	if !isHappy(1) {
		t.Error("error")
	}

	if !isHappy(19) {
		t.Error("error")
	}
	t.Log("success")
}
