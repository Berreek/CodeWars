package validParentheses_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/validParentheses"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	input        string
	output bool
}{
	{"()", true},
	{")", false},
	{")(()))", false},
	{"(", false},
	{"(())((()())())", true},
}

func TestValidParentheses(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value %v", i, tt.input), func(t *testing.T) {
			got := validParentheses.ValidParentheses(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}