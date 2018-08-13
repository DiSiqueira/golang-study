package ptypes_test

import (
	"github.com/disiqueira/golang-study/ex/ptypes"
	"testing"
)

func TestRightPropagate(t *testing.T) {
	for _, test := range []struct {
		in   int
		want int
	}{
		{128, 255},
		{64, 127},
		{36, 39},
		{1, 1},
		{2, 3},
		{3, 3},
	} {
		if got := ptypes.RightPropagate(test.in); got != test.want {
			t.Errorf("%.64b = %d; want %d", test.in, got, test.want)
		}
	}
}
