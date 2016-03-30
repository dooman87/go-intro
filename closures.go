package intro

func Sum(a func() int, b func() int) int {
	return a() + b()
}
