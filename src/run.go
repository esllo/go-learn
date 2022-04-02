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

func article5_Conditional() {
	// Statements

	// If
	num := 5
	if num == 3 {
		println("num is 3")
	} else if num == 5 {
		println("num is 5")
	} else {
		println("other")
	}

	// Optional
	if PI := 3.14; PI == 3.14 {
		println("PI = 3.14")
	}

	// Switch (always break)
	category := 1
	switch category {
	case 1:
		println("category 1")
	case 2:
		println("category 2")
	default:
		println("category 0")
	}

	// With Optional, Expression and fallthrough
	switch queue := 2; queue - 1 {
	case 1:
		println("queue 1")
		fallthrough // check next case
	case 2, 3: // multiple cases 2, 3
		println("queue 2")
		fallthrough // check next case
	default:
		println("queue end")
	}

	// Empty Switch => alternative (if... else if...else if...)
	name := "Park"
	switch {
	case name == "Kim":
		println("Kim")
	case name == "Park":
		println("park")
	case name == "Lee":
		println("Lee")
	}

	// Type Check Switch
	var flag interface{} = false
	switch flag.(type) {
	case int:
		println("flag is int")
	case string:
		println("flag is string")
	case bool:
		println("flag is bool")
	}
}

func main() {
	article(1, article1_Print)
	article(2, article2_FormattedPrint)
	article(3, article3_Variable)
	article(4, article4_DataType)
	article(5, article5_Conditional)
}
