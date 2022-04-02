package article

import "fmt"

func Article_Map() {
	var map1 map[int]string = nil
	map1 = map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
	}
	fmt.Printf("%v\n", map1)

	map2 := make(map[string]string)
	map2["Red"] = "255,0,0"
	map2["Blue"] = "0,0,255"
	map2["Green"] = "0, 255, 0"
	fmt.Printf("%v\n", map2)

	// check map key
	_, existsRed := map2["Red"]
	println("Red exists", existsRed)
	_, existsYellow := map2["Yellow"]
	println("Yellow exists", existsYellow)

	// for loop
	for key, value := range map2 {
		fmt.Printf("[%s] is (%s)\n", key, value)
	}
}
