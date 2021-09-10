package integersRecreation_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/integersRecreation"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	m int
	n int
	output [][]int
}{
	{1, 250, [][]int{{1, 1}, {42, 2500}, {246, 84100}}},
	{250, 500, [][]int{{287, 84100}}},
	{300, 600, [][]int{}},
}

func TestValidate(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value %v, %v", i, tt.m, tt.n), func(t *testing.T) {
			got := integersRecreation.ListSquared(tt.m, tt.n)
			require.Equal(t, got, tt.output)
		})
	}
}