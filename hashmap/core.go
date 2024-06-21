package hashmap

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
