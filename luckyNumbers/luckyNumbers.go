package luckyNumbers

// LuckyNumbers https://leetcode.com/problems/lucky-numbers-in-a-matrix/
func LuckyNumbers(matrix [][]int) []int {
	var minimums []int
	for i := 0; i < len(matrix); i++ {
		min := matrix[i][0]
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
			}
		}
		minimums = append(minimums, min)
	}

	var maximums []int
	for i := 0; i < len(matrix[0]); i++ {
		max := matrix[0][i]
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] > max {
				max = matrix[j][i]
			}
		}
		maximums = append(maximums, max)
	}

	var result []int
	for _, min := range minimums {
		for _, max := range maximums {
			if min == max {
				result = append(result, min)
			}
		}
	}

	return result
}
