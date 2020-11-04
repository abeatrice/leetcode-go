// https://leetcode.com/problems/shuffle-the-array/
// Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
// Return the array in the form [x1,y1,x2,y2,...,xn,yn].
package shuffleTheArray

func ShuffleTheArray(nums []int, n int) []int {
	var result []int
	for i := 0; i < len(nums)/2; i++ {
		result = append(result, nums[i])
		result = append(result, nums[i+n])
	}
	return result
}
