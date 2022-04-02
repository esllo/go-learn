package article

import "fmt"

func Article4_DataType() {
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
