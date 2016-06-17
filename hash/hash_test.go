package hash

import (
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
