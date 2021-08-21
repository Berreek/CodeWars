package createPhoneNumber_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/createPhoneNumber"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	input        [10]uint
	output string
}{
	{[10]uint {1,2,3,4,5,6,7,8,9,0}, "(123) 456-7890"},
}

func TestOddInt(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value %v", i, tt.input), func(t *testing.T) {
			got := createPhoneNumber.CreatePhoneNumber(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}