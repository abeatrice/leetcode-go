package kidsWithCandies_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/kidsWithCandies"
)

type testItem struct {
	candies      []int
	extraCandies int
	expected     []bool
}

func TestKidsWithCandies(t *testing.T) {
	testItems := []testItem{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	}

	for _, testItem := range testItems {
		results := kidsWithCandies.KidsWithCandies(testItem.candies, testItem.extraCandies)
		if len(results) != len(testItem.expected) {
			t.Errorf("kidsWithCandies(%v, %v), expected: %v, result: %v", testItem.candies, testItem.extraCandies, testItem.expected, results)
			continue
		}
		for i, result := range results {
			if testItem.expected[i] != result {
				t.Errorf("kidsWithCandies(%v, %v), expected: %v, result: %v", testItem.candies, testItem.extraCandies, testItem.expected, results)
				break
			}
		}
	}
}
