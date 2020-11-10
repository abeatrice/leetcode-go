package subtractProductAndSum

import (
	"strconv"
)

// https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
// Given an integer number n,
// return the difference between the product of its digits and the sum of its digits.
func SubtractProductAndSum(n int) int {
	numbers := strconv.Itoa(n)
	product := 1
	sum := 0
	for _, number := range numbers {
		int, _ := strconv.Atoi(string(number))
		product *= int
		sum += int
	}
	return product - sum
}
