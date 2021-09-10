package squares_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/squares"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	input int64
	output []int64
}{
	{50, []int64 {1,3,5,8,49}},
	{5, []int64 {3,4}},
	{2, []int64 {}},
	{7, []int64 {2,3,6}},
}

func TestValidate(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value %v", i, tt.input), func(t *testing.T) {
			got := squares.Decompose(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}