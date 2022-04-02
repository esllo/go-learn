package article

func Article5_Conditional() {
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
