package learningGame_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/learningGame"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	cmd    int
	input  []int
	output []int
}{
	{3, []int{4, 7, 10, 43, 1}, []int{400, 700, 1000, 4300, 100}},
	{1, []int{4, 7, 10, 43, 1}, []int{5, 8, 11, 44, 2}},
	{2, []int{8, 7, 10, 43, 4}, []int{0, 0, 0, 0, 0}},
}

func TestMachine_Command(t *testing.T) {
	sut := learningGame.NewMachine()

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for cmd: %v, input: %v, output %v", i, tt.cmd, tt.input, tt.output), func(t *testing.T) {
			for i, v := range tt.input {
				got := sut.Command(tt.cmd, v)
				sut.Response(got == tt.output[i])
			}

			for i, v := range tt.input {
				got := sut.Command(tt.cmd, v)
				require.Equal(t, got, tt.output[i])
			}
		})
	}
}
