package intro

func sum(a func() int, b func() int) int {
	return a() + b()
}

func DoSum() int {
	one := func() int {
		return 1
	}

	return sum(one, one)
}
