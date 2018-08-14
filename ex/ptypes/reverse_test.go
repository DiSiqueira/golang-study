package ptypes_test

import (
	"github.com/disiqueira/golang-study/ex/ptypes"
	"math"
	"math/bits"
	"math/rand"
	"testing"
)

func TestReverse(t *testing.T) {
	for i := 0; i < 10; i++ {
		num := uint64(rand.Intn(math.MaxInt32))
		want := bits.Reverse(uint(num))

		got := ptypes.Reverse(num)

		if got != uint64(want) {
			t.Errorf("Reverse(%d(%064b)) = %064b but want %064b", num, num, got, want)
		}
	}
}
