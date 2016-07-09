package math

func addDigits(num int) int {
	if num == 0 {
		return 0
	}

	i := num % 9
	if i == 0 {
		return 9
	}
	return i
}
