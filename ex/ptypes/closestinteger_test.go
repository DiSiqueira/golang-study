package ptypes_test

import (
	"github.com/disiqueira/golang-study/ex/ptypes"
	"testing"
)

func TestClosestInteger(t *testing.T) {
	for _, test := range []struct {
		num  uint64
		want uint64
	}{
		{6, 5},
		{65, 66},
		{76, 74},
	} {
		if got := ptypes.ClosestInteger(test.num); got != test.want {
			t.Errorf("Closest(%064b) = %064b , but want = %064b", test.num, got, test.want)
		}
	}
}
