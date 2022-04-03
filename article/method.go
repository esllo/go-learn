package article

import "fmt"

type Shape struct {
	width, height int
}

// value receiver
func (shape Shape) area() int {
	return shape.width * shape.height
}

// pointer receiver
func (shape *Shape) updateWidth(width int) int {
	shape.width = width
	return shape.width
}

func Article_Method() {
	shape := Shape{50, 10}
	fmt.Printf("shape : %d x %d = %d\n", shape.width, shape.height, shape.area())

	shape2 := Shape{20, 10}
	fmt.Printf("shape2 : %d x %d = %d\n", shape2.width, shape2.height, shape2.area())
	fmt.Println("update shape2 width to 30")
	shape2.updateWidth(30)
	fmt.Printf("shape2 : %d x %d = %d\n", shape2.width, shape2.height, shape2.area())
}
