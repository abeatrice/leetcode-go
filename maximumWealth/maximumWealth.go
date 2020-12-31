package maximumwealth

// MaximumWealth returns the richest customer wealth
// https://leetcode.com/problems/richest-customer-wealth/
func MaximumWealth(accounts [][]int) int {
	var result int
	var sum int
	for i := 0; i < len(accounts); i++ {
		sum = 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if sum > result {
			result = sum
		}
	}
	return result
}
