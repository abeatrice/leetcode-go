package numIdenticalPairs_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/numIdenticalPairs"
)

type testItem struct {
	nums     []int
	expected int
}

func TestNumIdenticalPairs(t *testing.T) {
	testItems := []testItem{
		{nums: []int{1, 2, 3, 1, 1, 3}, expected: 4},
		{nums: []int{1, 1, 1, 1}, expected: 6},
		{nums: []int{1, 2, 3}, expected: 0},
	}
	for _, testItem := range testItems {
		result := numIdenticalPairs.NumIdenticalPairs(testItem.nums)
		if result != testItem.expected {
			t.Errorf("NumIdenticalPairs(%v) Expected: %v, Result: %v", testItem.nums, testItem.expected, result)
		}
	}
}
