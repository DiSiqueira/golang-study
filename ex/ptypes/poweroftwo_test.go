package ptypes_test

import (
	"github.com/disiqueira/golang-study/ex/ptypes"
	"testing"
)

func TestPowerOfTwo(t *testing.T) {
	for _, test := range []struct {
		val  int
		want bool
	}{
		{1, true},
		{2, true},
		{4, true},
		{8, true},
		{128, true},
		{1024, true},
		{1 << 32, true},
		{5, false},
		{13, false},
		{1532, false},
	} {
		if got := ptypes.PowerOfTwo(test.val); got != test.want {
			t.Errorf("%d got = %t; want %t", test.val, got, test.want)
		}
	}
}
