package ptypes

func Swap(num uint64, i, j uint8) uint64 {
	if ((num >> i) & 1) != ((num >> j) & 1) {
		mask := (uint64(1) << i) | (uint64(1) << j)
		return num ^ mask
	}
	return num
}
