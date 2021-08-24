package bookseller_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/bookseller"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	books  []string
	categories  []string
	output string
}{
	{[]string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"},  []string{"A", "B", "C", "D"}, "(A : 0) - (B : 1290) - (C : 515) - (D : 600)"},
	{[]string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"},  []string{"A", "B"},  "(A : 200) - (B : 1140)"},
	{[]string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"},  []string{"P", "U"},  "(P : 0) - (U : 0)"},
	{[]string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"},  []string{},  ""},
}

func TestStockList(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value %s", i, tt.books), func(t *testing.T) {
			got := bookseller.StockList(tt.books, tt.categories)
			require.Equal(t, got, tt.output)
		})
	}
}