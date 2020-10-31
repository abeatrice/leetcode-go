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
