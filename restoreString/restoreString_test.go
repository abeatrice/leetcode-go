package restoreString_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/restoreString"
)

type testItem struct {
	s        string
	indices  []int
	expected string
}

func TestRestoreString(t *testing.T) {
	testItems := []testItem{
		{s: "codeleet", indices: []int{4, 5, 6, 7, 0, 2, 1, 3}, expected: "leetcode"},
		{s: "abc", indices: []int{0, 1, 2}, expected: "abc"},
		{s: "aiohn", indices: []int{3, 1, 4, 2, 0}, expected: "nihao"},
		{s: "aaiougrt", indices: []int{4, 0, 2, 6, 7, 3, 1, 5}, expected: "arigatou"},
		{s: "art", indices: []int{1, 0, 2}, expected: "rat"},
	}
	for _, testItem := range testItems {
		result := restoreString.RestoreString(testItem.s, testItem.indices)
		if result != testItem.expected {
			t.Errorf("RestoreString(%v, %v) Result: %v, Expected %v", testItem.s, testItem.indices, result, testItem.expected)
		}
	}
}
