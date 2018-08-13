package ptypes_test

import (
	"github.com/disiqueira/golang-study/ex/ptypes"
	"testing"
)

const (
	odd  = 1
	even = 0
)

func TestParity(t *testing.T) {
	for _, test := range []struct {
		in   uint64
		want uint8
	}{
		{0, even},
		{1, odd},
		{2, odd},
		{3, even},
		{4, odd},
		{5, even},
		{6, even},
		{7, odd},
		{8, odd},
		{9, even},
		{1<<64 - 1, even},
	} {
		if got := ptypes.Parity(test.in); got != test.want {
			t.Errorf("%.64b = %d; want %d", test.in, got, test.want)
		}
	}
}

func BenchmarkParity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ptypes.Parity(1<<64 - 1)
	}
}
