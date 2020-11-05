package numJewelsInStones_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/numJewelsInStones"
)

type testItem struct {
	jewels   string
	stones   string
	expected int
}

func TestNumJewelsInStones(t *testing.T) {
	testItems := []testItem{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}
	for _, testItem := range testItems {
		result := numJewelsInStones.NumJewelsInStones(testItem.jewels, testItem.stones)
		if result != testItem.expected {
			t.Errorf("NumJewelsInStones(%v, %v) expected: %v, result %v", testItem.jewels, testItem.stones, testItem.expected, result)
		}
	}
}
