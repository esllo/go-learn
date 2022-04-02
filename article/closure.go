package article

func incrementClosure() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}

func name(first string) func(string) string {
	return func(second string) string {
		return first + " " + second
	}
}

func Article_Closure() {
	var increment func() int = incrementClosure()
	println(increment())
	println(increment())
	println(increment())

	hong := name("Hong")
	println(hong("Gildong"))
	println(hong("Cheolsoo"))
}
