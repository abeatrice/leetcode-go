package maximumWealth_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/maximumWealth"
)

type testItem struct {
	accounts [][]int
	expected int
}

func TestMaximumWealth(t *testing.T) {
	testItems := []testItem{
		{accounts: [][]int{{1, 2, 3}, {3, 2, 1}}, expected: 6},
		{accounts: [][]int{{1, 5}, {7, 3}, {3, 5}}, expected: 10},
		{accounts: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, expected: 17},
	}
	for _, testItem := range testItems {
		result := maximumWealth.MaximumWealth(testItem.accounts)
		if result != testItem.expected {
			t.Errorf("MaximumWealth(%v) Expected: %v, Result: %v", testItem.accounts, testItem.expected, result)
		}
	}
}
