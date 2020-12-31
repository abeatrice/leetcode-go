package interpret_test

import (
	"testing"

	"github.com/abeatrice/leetcode-go/interpret"
)

type testItem struct {
	command  string
	expected string
}

func TestInterpret(t *testing.T) {
	testItems := []testItem{
		{command: "G()(al)", expected: "Goal"},
		{command: "G()()()()(al)", expected: "Gooooal"},
		{command: "(al)G(al)()()G", expected: "alGalooG"},
	}
	for _, testItem := range testItems {
		result := interpret.Interpret(testItem.command)
		if result != testItem.expected {
			t.Errorf("Interpret(%v) Expected: %v, Result: %v", testItem.command, testItem.expected, result)
		}
	}
}
