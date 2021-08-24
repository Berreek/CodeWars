package outlier_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/outlier"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

var tests = []struct {
	input        []int
	output int
}{
	{[]int {2, 6, 8, -10, 3}, 3},
	{[]int {206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781}, 206847684},
	{[]int {math.MaxInt32, 0, 1}, 0},
}

func TestFindOutlier(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value %v", i, tt.input), func(t *testing.T) {
			got := outlier.FindOutlier(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}
