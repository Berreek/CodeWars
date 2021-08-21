package countDuplicates_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/countDuplicates"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	input string
	output int
}{
	{"abcde", 0},
	{"aabbcde", 2},
	{"aabBcde", 2},
	{"indivisibility", 1},
	{"Indivisibilities", 2},
	{"aA11", 2},
	{"ABBA", 2},
}

func TestDuplicateCount(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value %s", i, tt.input), func(t *testing.T) {
			got := countDuplicates.DuplicateCount(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}