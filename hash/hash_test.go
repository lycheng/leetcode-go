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

func TestCode36(t *testing.T) {
	board := [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}

	if !isValidSudoku(board) {
		t.Error("error")
	}

	board1 := [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.1..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	if isValidSudoku(board1) {
		t.Error("error")
	}
	t.Log("success")
}

func TestCode205(t *testing.T) {
	trueTuple := [][]string{
		[]string{"egg", "add"},
		[]string{"paper", "title"},
	}

	for _, items := range trueTuple {
		if !isIsomorphic(items[0], items[1]) {
			t.Error("error")
		}
	}

	falseTuple := [][]string{
		[]string{"foo", "bar"},
		[]string{"ab", "aa"},
	}

	for _, items := range falseTuple {
		if isIsomorphic(items[0], items[1]) {
			t.Error("error")
		}
	}
	t.Log("success")
}

func TestCode347(t *testing.T) {
	src := []int{1, 1, 1, 2, 2, 3}
	target := []int{1, 2}

	if !reflect.DeepEqual(target, topKFrequent(src, len(target))) {
		t.Error("error")
	}
	t.Log("success")
}

func TestCode500(t *testing.T) {
	src := []string{"Hello", "Alaska", "Dad", "Peace"}
	dst := []string{"Alaska", "Dad"}

	if ret := findWords(src); !reflect.DeepEqual(dst, ret) {
		t.Fatal("error")
	}
	t.Log("success")
}

func TestCode575(t *testing.T) {
	candies := []int{1, 1, 3, 4}
	if k := distributeCandies(candies); k != 2 {
		t.Fatal("expect 2 and got", k)
	}
	candies = []int{1000, 1000, 2, 1, 2, 5, 3, 1}
	if k := distributeCandies(candies); k != 4 {
		t.Fatal("expect 4 and got", k)
	}
	t.Log("success")
}
