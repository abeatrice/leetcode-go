package luckyNumbers_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/helper"
	"github.com/abeatrice/leetcode-go/luckyNumbers"
)

type testItem struct {
	matrix   [][]int
	expected []int
}

func TestLuckyNumbers(t *testing.T) {
	testItems := []testItem{
		{matrix: [][]int{[]int{3, 7, 8}, []int{9, 11, 13}, []int{15, 16, 17}}, expected: []int{15}},
		{matrix: [][]int{[]int{1, 10, 4, 2}, []int{9, 3, 8, 7}, []int{15, 16, 17, 12}}, expected: []int{12}},
		{matrix: [][]int{[]int{7, 8}, []int{1, 2}}, expected: []int{7}},
	}
	for _, testItem := range testItems {
		result := luckyNumbers.LuckyNumbers(testItem.matrix)
		if !helper.ArrayIsEqual(result, testItem.expected) {
			t.Errorf("LuckyNumbers(%v) Expected: %v, Result: %v", testItem.matrix, testItem.expected, result)
		}
	}
}
