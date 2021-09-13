package friendCheating_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/friendCheating"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	input  uint64
	output [][2]uint64
}{
	{26, [][2]uint64{{15, 21}, {21, 15}}},
	{100, nil},
	{101, [][2]uint64{{55, 91}, {91, 55}}},
	{102, [][2]uint64{{70, 73}, {73, 70}}},
	{1000003, [][2]uint64{{550320, 908566}, {559756, 893250}, {893250, 559756}, {908566, 550320}}},
}

func TestRemoveNb(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value %v", i, tt.input), func(t *testing.T) {
			got := friendCheating.RemovNb(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}
