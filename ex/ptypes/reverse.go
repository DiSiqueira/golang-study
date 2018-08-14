package ptypes

func Reverse(num uint64) uint64 {
	res := uint64(0)
	for i := 0; i < 64; i++ {
		res <<= 1
		res |= num & uint64(1)
		num >>= 1
	}
	return res
}
