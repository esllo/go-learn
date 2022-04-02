package article

func Article6_Loop() {
	// for
	for i := 0; i < 3; i++ {
		print(i, " ")
	}
	println()

	// for (like while)
	i := 0
	for i < 3 {
		print(i, " ")
		i++
	}
	println()

	// for infinite
	j := 0
	for {
		print(j, " ")
		j++
		if j == 3 {
			break
		}
	}
	println()

	// for range (foreach)
	scores := []int{90, 80, 70}
	// _(underscore) is Blank Identifier
	for _, score := range scores {
		print(score, " ")
	}
	println()

	// for continue, goto, break
	nums := []int{1, 2, 3, 4, 5}
LOOP:
	for _, num := range nums {
		if num == 2 {
			continue
		}
		print(num, " ")
		if num == 4 {
			goto HERE
		}
		if num == 5 {
			break LOOP
		}
	}
HERE:
	println("come HERE")
}
