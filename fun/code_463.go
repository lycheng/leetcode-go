package fun

// https://leetcode.com/problems/island-perimeter/description/

func islandPerimeter(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			result += perimeter(grid, i, j, len(grid), len(grid[i]))
		}
	}
	return result
}

func perimeter(grid [][]int, i, j int, maxI, maxJ int) int {
	total := 4
	if i != 0 && grid[i-1][j] == 1 {
		total -= 1
	}

	if i < maxI-1 && grid[i+1][j] == 1 {
		total -= 1
	}

	if j != 0 && grid[i][j-1] == 1 {
		total -= 1
	}

	if j < maxJ-1 && grid[i][j+1] == 1 {
		total -= 1
	}
	return total
}
