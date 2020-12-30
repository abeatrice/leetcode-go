// https://leetcode.com/problems/richest-customer-wealth/
// You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest customer has.
// A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.
package maximumWealth

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
