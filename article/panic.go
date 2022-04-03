package article

import (
	"fmt"
	"os"
)

func openFile(fn string) {
	defer func() {
		// recover panic
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		// make panic
		panic(err)
	}

	defer f.Close()
}

func checkErrorTypeAssertion() {
	defer func() {
		err := recover()
		// error || *CustomError
		if _, ok := err.(*CustomError); ok {
			fmt.Println(err)
		}
	}()

	_, err := requestError()
	if err != nil {
		panic(err)
	}
}

func Article_Panic() {
	// Invalid.txt not exists
	openFile("Invalid.txt")

	println("Done, Recover")

	checkErrorTypeAssertion()

	println("Done, Type Assertion Check")
}
