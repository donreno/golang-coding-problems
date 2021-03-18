package googleinterview_test

import (
	gi "golang-coding-problems/internal/googleinterview"

	"fmt"
	"testing"

	"github.com/franela/goblin"
)

func TestIsValidNumber(t *testing.T) {
	g := goblin.Goblin(t)

	testCases := []struct {
		num   string
		valid bool
	}{
		{num: "4", valid: true},
		{num: "43.50", valid: true},
		{num: "43", valid: true},
		{num: "43.50.", valid: false},
		{num: "4X", valid: false},
	}

	g.Describe("Is Valid Number", func() {

		for _, testCase := range testCases {
			g.It(fmt.Sprintf("When input is %q then result should be %v", testCase.num, testCase.valid), func() {
				g.Assert(gi.IsValidNumber(testCase.num)).Eql(testCase.valid)
			})
		}
	})
}
