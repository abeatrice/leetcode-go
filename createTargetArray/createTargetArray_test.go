package createTargetArray_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/createTargetArray"
)

type testItem struct {
	nums     []int
	index    []int
	expected []int
}

func TestCreateTargetArray(t *testing.T) {
	testItems := []testItem{
		{nums: []int{0, 1, 2, 3, 4}, index: []int{0, 1, 2, 2, 1}, expected: []int{0, 4, 1, 3, 2}},
		// {nums: []int{1, 2, 3, 4, 0}, index: []int{0, 1, 2, 3, 0}, expected: []int{0, 1, 2, 3, 4}},
		// {nums: []int{1}, index: []int{0}, expected: []int{1}},
	}
	for _, testItem := range testItems {
		result := createTargetArray.CreateTargetArray(testItem.nums, testItem.index)
		if len(result) != len(testItem.expected) {
			t.Errorf("CreateTargetArray(%v,%v) Expected: %v, Result, %v", testItem.nums, testItem.index, testItem.expected, result)
		} else if !equal(result, testItem.expected) {
			t.Errorf("CreateTargetArray(%v,%v) Expected: %v, Result, %v", testItem.nums, testItem.index, testItem.expected, result)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
