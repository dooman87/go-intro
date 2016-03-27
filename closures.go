package intro

func sum(a func() int, b func() int) int {
	return a() + b()
}

func DoSum() {
	one := func() int {
		return 1
	}

	sum(one, one)
}
