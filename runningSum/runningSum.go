// https://leetcode.com/problems/running-sum-of-1d-array/
// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
// Return the running sum of nums.
package runningSum

func RunningSum(nums []int) []int {
	var results []int
	var sum int
	for _, num := range nums {
		sum += num
		results = append(results, sum)
	}
	return results
}
