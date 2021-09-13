package weirdPrimeGenerator_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/weirdPrimeGenerator"
	"github.com/stretchr/testify/require"
	"testing"
)

var testsCountOnes = []struct {
	input int64
	output int
}{
	{10, 8},
	{100, 90},
}

func TestCountOnes(t *testing.T) {
	for i, tt := range testsCountOnes {
		t.Run(fmt.Sprintf("Case %d for value %v", i, tt.input), func(t *testing.T) {
			got := weirdPrimeGenerator.CountOnes(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}

var testsMaxPn = []struct {
	input int64
	output int64
}{
	{5, 47},
	{7, 101},
}

func TestMaxPn(t *testing.T) {
	for i, tt := range testsMaxPn {
		t.Run(fmt.Sprintf("Case %d for value %v", i, tt.input), func(t *testing.T) {
			got := weirdPrimeGenerator.MaxPn(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}


var testsAnOverAverage = []struct {
	input int64
	output int
}{
	{5, 3},
}

func TestAnOverAverage(t *testing.T) {
	for i, tt := range testsAnOverAverage {
		t.Run(fmt.Sprintf("Case %d for value %v", i, tt.input), func(t *testing.T) {
			got := weirdPrimeGenerator.AnOverAverage(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}