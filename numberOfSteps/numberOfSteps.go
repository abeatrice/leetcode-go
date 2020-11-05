package numberOfSteps

func NumberOfSteps(num int) int {
	var steps int
	for num != 0 {
		steps++
		if num%2 == 0 {
			num = num / 2
		} else {
			num--
		}
	}
	return steps
}
