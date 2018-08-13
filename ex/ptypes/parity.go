package ptypes

func Parity(num uint64) uint8 {
	p := uint8(0)
	for num > 0 {
		p ^= 1
		num &= num - 1
	}
	return p
}
