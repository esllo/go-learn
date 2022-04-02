package article

func printString(msg string) {
	println("prt str : ", msg)
}

func add(a int, b int) int {
	return a + b
}

func callByReference(msg *string) {
	*msg = "reference"
}

func variadicParams(nums ...int) (int, int) {
	sum := 0
	count := 0
	for _, num := range nums {
		sum += num
		count++
	}
	return sum, count
}

func namedReturnParams(nums ...int) (sum int, count int) {
	// sum = 0, count = 0
	for _, num := range nums {
		sum += num
		count++
	}
	return
}

func Article_Function() {
	printString("hi")
	println(add(3, 5))
	var msg string = "value"
	println("origin : ", msg)
	callByReference(&msg)
	println("changed : ", msg)
	variadicParams(1, 2, 3)
	namedReturnParams(1, 2, 3)
}
