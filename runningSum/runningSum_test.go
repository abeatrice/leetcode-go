package runningSum_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/runningSum"
)

type TestItem struct {
	input    []int
	expected []int
}

func TestRunningSum(t *testing.T) {
	items := []TestItem{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}

	for _, item := range items {
		sums := runningSum.RunningSum(item.input)
		for i, _ := range sums {
			if sums[i] != item.expected[i] {
				t.Errorf("RunningSum() expected: %v, result: %v \n", item.expected, sums)
				break
			}
		}
	}
}
