package countConsistentStrings_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/countConsistentStrings"
)

type testItem struct {
	allowed  string
	words    []string
	expected int
}

func TestCountConsistentStrings(t *testing.T) {
	testItems := []testItem{
		{allowed: "ab", words: []string{"ad", "bd", "aaab", "baa", "badab"}, expected: 2},
		{allowed: "abc", words: []string{"a", "b", "c", "ab", "ac", "bc", "abc"}, expected: 7},
		{allowed: "cad", words: []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}, expected: 4},
	}
	for _, testItem := range testItems {
		result := countConsistentStrings.CountConsistentStrings(testItem.allowed, testItem.words)
		if result != testItem.expected {
			t.Errorf("CountConsistentStrings(%v,%v) Expected: %v, Result: %v", testItem.allowed, testItem.words, testItem.expected, result)
		}
	}
}
