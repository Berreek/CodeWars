package alphanumeric_test

import (
	"fmt"
	"github.com/Berreek/CodeWars/alphanumeric"
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	input string
	output bool
}{
	{".*?", false},
	{"a", true},
	{"Mazinkaiser", true},
	{"hello world_", false},
	{"PassW0rd", true},
	{"          ", false},
	{"", false},
	{"\\n\\t\\n", false},
	{"ciao\\n$$_", false},
	{"__ * __", false},
	{"&)))(((", false},
	{"43534h56jmTHHF3k", true},
}

func TestValidate(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d for value %s", i, tt.input), func(t *testing.T) {
			got := alphanumeric.Alphanumeric(tt.input)
			require.Equal(t, got, tt.output)
		})
	}
}