package sumByFactors_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/sumByFactors"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	input  []int
	output string
}{
	{[]int{12, 15}, "(2 12)(3 27)(5 15)"},
	{[]int{15,21,24,30,45}, "(2 54)(3 135)(5 90)(7 21)"},
	{[]int{114, 237, 421}, "(2 114)(3 351)(19 114)(79 237)(421 421)"},
}

func TestSumOfDivided(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value %v", i, tt.input), func(t *testing.T) {
			got := sumByFactors.SumOfDivided(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}
