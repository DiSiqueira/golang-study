package ptypes_test

import (
	"github.com/disiqueira/golang-study/ex/ptypes"
	"testing"
)

func TestModPowerOfTwo(t *testing.T) {
	for _, test := range []struct {
		x, y int
		want int
	}{
		{64, 77, 13},
		{64, 70, 6},
		{2, 4, 0},
		{2, 5, 1},
	} {
		if got := ptypes.ModPowerOfTwo(test.x, test.y); got != test.want {
			t.Errorf("%d MOD %d = %d; want %d", test.x, test.y, got, test.want)
		}
	}
}
