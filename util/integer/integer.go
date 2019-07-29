package integer

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	return a - Max(a, b) + b // a + b - Max(a, b)
}
