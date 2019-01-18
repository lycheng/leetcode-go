package array

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] == target {
			break
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{i + 1, j + 1}
}
