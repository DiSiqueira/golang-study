package ptypes_test

import (
	"github.com/disiqueira/golang-study/ex/ptypes"
	"testing"
)

func TestSwap(t *testing.T) {
	for _, test := range []struct {
		num  uint64
		i, j uint8
		want uint64
	}{
		{73, 6, 1, 11},
		{1, 0, 1, 1 << 1},
		{5, 2, 4, 17},
		{20, 2, 4, 20},
		{1, 0, 63, 1 << 63},
	} {
		if got := ptypes.Swap(test.num, test.i, test.j); got != test.want {
			t.Errorf("%b swap %d and %d: got: %b want: %b", test.num, test.i, test.j, got, test.want)
		}
	}
}
