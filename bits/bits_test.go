// package bits for leetcode bit-manipulation
package bits

import (
	"math/rand"
	"testing"
	"time"
)

func TestCode461(t *testing.T) {

	rand.Seed(time.Now().Unix())

	x := rand.Int()
	y := rand.Int()

	x %= 2147483647
	y %= 2147483647
	if hammingDistance(x, y) != hammingDistance1(x, y) {
		t.Errorf("Error @ (%d %d)", x, y)
	}
	t.Log("Success")
}

func TestCode477(t *testing.T) {

	nums := []int{4, 14, 2}

	if totalHammingDistance(nums) != 6 {
		t.Error("Error")
	}
	t.Log("Success")
}
