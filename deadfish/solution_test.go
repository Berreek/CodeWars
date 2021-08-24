package deadfish_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/deadfish"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	input        string
	output []int
}{
	{"ooo", []int{0,0,0}},
	{"ioioio", []int{1,2,3}},
	{"idoiido", []int{0,1}},
	{"isoisoiso", []int{1,4,25}},
	{"codewars", []int{0}},
	{"ssdizbsnsdkkincuybidds", make([]int, 0)},
}

func TestParse(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value %v", i, tt.input), func(t *testing.T) {
			got := deadfish.Parse(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}