package ptypes

func RightPropagate(num int) int {
	return num | (num - 1)
}
