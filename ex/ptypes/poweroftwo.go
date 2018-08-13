package ptypes

func PowerOfTwo(num int) bool {
	return num&(num-1) == 0
}
