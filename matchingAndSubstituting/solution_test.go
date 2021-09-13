package matchingAndSubstituting_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/matchingAndSubstituting"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	s  string
	prog  string
	version  string
	output string
}{
	{"Program title: Primes\nAuthor: Kern\nCorporation: Gold\nPhone: +1-503-555-0091\nDate: Tues April 9, 2005\nVersion: 6.7\nLevel: Alpha","Ladder", "1.1", "Program: Ladder Author: g964 Phone: +1-503-555-0090 Date: 2019-01-01 Version: 1.1"},
	{"Program title: Hammer\nAuthor: Tolkien\nCorporation: IB\nPhone: +1-503-555-0090\nDate: Tues March 29, 2017\nVersion: 2.0\nLevel: Release", "Balance", "1.5.6", "Program: Balance Author: g964 Phone: +1-503-555-0090 Date: 2019-01-01 Version: 2.0"},
	{"Program title: Primes\nAuthor: Kern\nCorporation: Gold\nPhone: +1-503-555-009\nDate: Tues April 9, 2005\nVersion: 6.7\nLevel: Alpha", "Ladder", "1.1", "ERROR: VERSION or PHONE"},
	}

func TestChange(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value s: %s, prog: %s, version: %s", i, tt.s, tt.prog, tt.version), func(t *testing.T) {
			got := matchingAndSubstituting.Change(tt.s, tt.prog, tt.version)
			require.Equal(t, got, tt.output)
		})
	}
}