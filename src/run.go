package main

import (
	"fmt"
)

func article(i int, f func()) {
	println("================")
	println("article", i)
	println("----------------")
	f()
	println("\n================\n\n")
}

func article1_Print() {
	print("Hello")
	println(" world!")
}

type Data struct {
	one   int
	two   string
	three bool
}

func article2_FormattedPrint() {
	fmt.Print("Hello")
	fmt.Println(" world!")
	fmt.Printf("%d\n%f\n%s\n%t\n%v\n%p\n", 1, 1.0, "string", false, []int{}, new(int))
	fmt.Println()
	fmt.Printf("%d, %o, %b, %x, %X\n", 12, 12, 12, 12, 12)
	fmt.Println()
	var data = Data{1, "two", false}
	fmt.Printf("%v\n%+v\n%#v", data, data, data)
}

func article3_Variable() {
	var numImplicit = 3
	var numExplicit int32 = 3

	fmt.Printf("Imp %d, Exp %d\n", numImplicit, numExplicit)

	const numConstant = 3.14
	const myaConstant = "먀"
	fmt.Printf("NC %f, SC %s\n", numConstant, myaConstant)
}

func article4_DataType() {
	// Types
	/*
		int int8 int16 int32 int64
		uint uint8 ... uintptr
		float32 float64 complex64 complex128
		bool
		string(immutable)
			`` => Raw String Literal(Multiline) > \n => \n
			"" => Interpreted String Literal > \n => New Line
		byte(8)
		rune(32) Unicode Code point ex) 00000430 = A
	*/
	// Cast
	var numInt int = 100
	var numUInt = uint(numInt)
	var numFloat = float32(numUInt)
	fmt.Printf("%d, %d, %f\n", numInt, numUInt, numFloat)

	// ":=" = Short Assignment Statement (define and assignment)
	str := "Hello world"
	bytes := []byte(str)
	str2 := string(bytes)
	fmt.Printf("bytes : %s\nstr   : %s", bytes, str2)
}

func main() {
	article(1, article1_Print)
	article(2, article2_FormattedPrint)
	article(3, article3_Variable)
	article(4, article4_DataType)
}
