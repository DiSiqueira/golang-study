package ptypes

func ModPowerOfTwo(x, y int) int {
	return y & (x - 1)
}
