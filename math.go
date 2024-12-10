package batteries

func AbsInt(num int) int {
	if num < 0 {
		return 0 - num
	}

	return num
}
