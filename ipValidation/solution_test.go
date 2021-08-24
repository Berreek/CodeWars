package ipValidation_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/ipValidation"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	input        string
	output bool
}{
	{"12.255.56.1", true},
	{"", false},
	{"abc.def.ghi.jkl", false},
	{"123.456.789.0", false},
	{"12.34.56", false},
	{"12.34.56 .1", false},
	{"12.34.56.-1", false},
	{"123.045.067.089", false},
	{"127.1.1.0", true},
	{"0.0.0.0", true},
	{"0.34.82.53", true},
	{"192.168.1.300", false},
}

func TestIpValidation(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value %v", i, tt.input), func(t *testing.T) {
			got := ipValidation.Is_valid_ip(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}