package defangIPaddr_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/defangIPaddr"
)

type testItem struct {
	address  string
	expected string
}

func TestDefangIPaddr(t *testing.T) {
	testItems := []testItem{
		{"1.1.1.1", "1[.]1[.]1[.]1"},
		{"255.100.50.0", "255[.]100[.]50[.]0"},
	}
	for _, testItem := range testItems {
		result := defangIPaddr.DefangIPaddr(testItem.address)
		if result != testItem.expected {
			t.Errorf("defangIPaddr(%v) expected: %v, result: %v", testItem.address, testItem.expected, result)
		}
	}
}
