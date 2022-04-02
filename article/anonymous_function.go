package article

type calculartor func(int, int) int

func calc(f calculartor, a int, b int) (result int) {
	result = f(a, b)
	return
}

func Article_AnonymousFunction() {
	var add func(int, int) int = func(a int, b int) int {
		return a + b
	}
	println(add(3, 5))

	sub := func(x int, y int) int {
		return x - y
	}

	println(calc(add, 3, 5))
	println(calc(sub, 5, 3))
}
