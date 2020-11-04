package shuffleTheArray_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/shuffleTheArray"
)

type testItem struct {
	nums     []int
	n        int
	expected []int
}

func TestShufffleTheArray(t *testing.T) {
	testItems := []testItem{
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}

	for _, testItem := range testItems {
		results := shuffleTheArray.ShuffleTheArray(testItem.nums, testItem.n)
		if len(results) != len(testItem.expected) {
			t.Errorf("shuffleTheArray(%v, %v) expected: %v, result: %v", testItem.nums, testItem.n, testItem.expected, results)
			continue
		}
		for i, result := range results {
			if result != testItem.expected[i] {
				t.Errorf("shuffleTheArray(%v, %v) expected: %v, result: %v", testItem.nums, testItem.n, testItem.expected, results)
				break
			}
		}
	}
}
