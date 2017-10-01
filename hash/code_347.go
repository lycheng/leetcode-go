package hash

// https://leetcode.com/problems/top-k-frequent-elements/description/
func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int, 0)
	for _, num := range nums {
		counts[num] = counts[num] + 1
	}

	frequestList := make([][]int, len(nums)+1)
	for num, count := range counts {
		frequestList[count] = append(frequestList[count], num)
	}
	result := make([]int, 0)
	for j := len(nums); j >= 0; j-- {
		if len(frequestList[j]) == 0 {
			continue
		}
		result = append(result, frequestList[j]...)
		if len(result) >= k {
			break
		}
	}
	return result[:k]
}
