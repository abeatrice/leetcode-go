package numberOfSteps_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/numberOfSteps"
)

type testItem struct {
	num      int
	expected int
}

func TestNumberOfSteps(t *testing.T) {
	testItems := []testItem{
		{14, 6},
		{8, 4},
		{123, 12},
	}
	for _, testItem := range testItems {
		result := numberOfSteps.NumberOfSteps(testItem.num)
		if result != testItem.expected {
			t.Errorf("NumberOfSteps(%v) expected:%v, result:%v", testItem.num, testItem.expected, result)
		}
	}
}
