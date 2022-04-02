package article

import "fmt"

func Article3_Variables() {
	var numImplicit = 3
	var numExplicit int32 = 3

	fmt.Printf("Imp %d, Exp %d\n", numImplicit, numExplicit)

	const numConstant = 3.14
	const myaConstant = "먀"
	fmt.Printf("NC %f, SC %s\n", numConstant, myaConstant)
}
