package article

import "fmt"

func Article_Array() {
	// var array1 [3]int
	array1 := [3]int{1, 2, 3}
	array2 := [...]int{1, 2, 3}
	fmt.Printf("%v %v\n", array1, array2)

	// var array2d [2][2]int
	array2d := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Printf("%v\n", array2d)

	println("slice")
	var slice1 []int
	slice1 = []int{1, 2, 3}
	// make(type, length, capacity)
	slice2 := make([]int, 3, 6)
	fmt.Printf("slice1 leng=%d cap=%d %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2 leng=%d cap=%d %v\n", len(slice2), cap(slice2), slice2)

	// append
	println("append")
	slice1 = append(slice1, 4, 5, 6, 7)
	slice2 = append(slice2, 0, 0, 0)
	fmt.Printf("slice1 leng=%d cap=%d %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2 leng=%d cap=%d %v\n", len(slice2), cap(slice2), slice2)

	// copy
	println("copy")
	slice3 := make([]int, len(slice1), cap(slice1))
	copy(slice3, slice1)
	fmt.Printf("slice1 leng=%d cap=%d %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("slice3 leng=%d cap=%d %v\n", len(slice3), cap(slice3), slice3)

	// sub-slice
	println("sub-slice")
	sub := slice3[2:6]
	fmt.Printf("%v\n", sub)
}
