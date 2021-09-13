package smallestSum_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/smallestSum"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	input  []int
	output int
}{
	{[]int{9}, 9},
	{[]int{6, 9, 21}, 9},
	{[]int{1, 21, 55}, 3},
	{[]int{1, 21, 55}, 3},
	{[]int{3, 1, 2, 3, 1}, 5},
}

func TestSolution(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value %v", i, tt.input), func(t *testing.T) {
			got := smallestSum.Solution(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}
