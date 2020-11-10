package subtractProductAndSum_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/subtractProductAndSum"
)

type testItem struct {
	input    int
	expected int
}

func TestSubtractProductAndSum(t *testing.T) {
	testItems := []testItem{
		{234, 15},
		{4421, 21},
	}
	for _, testItem := range testItems {
		result := subtractProductAndSum.SubtractProductAndSum(testItem.input)
		if result != testItem.expected {
			t.Errorf("SubtractProductAndSum(%v) expected: %v, result: %v", testItem.input, testItem.expected, result)
		}
	}
}
