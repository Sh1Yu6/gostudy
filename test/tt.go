package test

func F(v int) func(int) int {
	sum := v
	return func(v int) int {
		sum += v
		return sum
	}
}
