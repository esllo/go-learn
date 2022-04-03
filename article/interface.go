package article

import "fmt"

// interface
type Phone interface {
	size() float64
}

// galaxy struct
type Galaxy struct {
	width, height float64
}

// iphone struct
type IPhone struct {
	width, height float64
}

// interface impl
func (galaxy *Galaxy) size() float64 {
	return galaxy.width * galaxy.height
}

// interface impl
func (iphone *IPhone) size() float64 {
	return iphone.width * iphone.height
}

// print interface
func printSize(phones ...Phone) {
	for index, phone := range phones {
		fmt.Printf("phone %d's size = %.2f\n", index, phone.size())
	}
}

func Article_Interface() {
	// interface implementation
	galaxy := Galaxy{70.6, 146.0}
	iphone := IPhone{71.5, 146.7}

	printSize(&galaxy, &iphone)

	// interface{} = Dynamic Type (like java Object)
	var x interface{} = nil
	x = 3
	fmt.Println("x =", x)
	x = "Text"
	fmt.Println("x =", x)
	x = galaxy
	fmt.Println("x =", x)

	// type asserion
	var y interface{} = 1

	i := y
	// type assertion
	// if y == nil || typeof y != int => Runtime Error
	j := y.(int)

	println(i) // print pointer address
	println(j) // print 1
}
