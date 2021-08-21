package oddInt_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/oddInt"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	input        []int
	output int
}{
	{[]int {20,1,-1,2,-2,3,3,5,5,1,2,4,20,4,-1,-2,5}, 5},
	{[]int {1,1,2,-2,5,2,4,4,-1,-2,5}, -1},
	{[]int {20,1,1,2,2,3,3,5,5,4,20,4,5}, 5},
	{[]int {10}, 10},
	{[]int {1,1,1,1,1,1,10,1,1,1,1}, 10},
	{[]int {5,4,3,2,1,5,4,3,2,10,10}, 1},
}

func TestFindOdd(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value %v", i, tt.input), func(t *testing.T) {
			got := oddInt.FindOdd(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}