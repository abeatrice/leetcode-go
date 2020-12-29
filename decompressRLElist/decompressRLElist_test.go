package decompressRLElist_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/decompressRLElist"
)

type testItem struct {
	nums     []int
	expected []int
}

func TestDecompressRLElist(t *testing.T) {
	testItems := []testItem{
		{nums: []int{1, 2, 3, 4}, expected: []int{2, 4, 4, 4}},
		{nums: []int{1, 1, 2, 3}, expected: []int{1, 3, 3}},
	}
	for _, testItem := range testItems {
		result := decompressRLElist.DecompressRLElist(testItem.nums)
		if !equal(result, testItem.expected) {
			t.Errorf("DecompressRLElist(%v), Expected: %v, Result:%v", testItem.nums, testItem.expected, result)
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
