package ptypes

func ClosestInteger(num uint64) uint64 {
	for i := uint64(0); i < 64; i++ {
		if ((num >> i) & 1) != ((num >> (i + 1)) & 1) {
			return num ^ (1<<i | 1<<(i+1))
		}
	}
	return 0
}
