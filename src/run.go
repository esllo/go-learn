package main

import "fmt"

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

func main() {
	article1_Print()
	article2_FormattedPrint()
}
