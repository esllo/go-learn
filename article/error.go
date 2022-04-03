package article

import (
	"fmt"
	"os"
)

type CustomError struct {
	Code int
}

func (customError *CustomError) Error() string {
	return fmt.Sprintf("error code : %d", customError.Code)
}

func requestError() (int, error) {
	return 404, &CustomError{
		Code: 404,
	}
}

func Article_Error() {
	f, err := os.Open("C:\\temp\\1.txt")
	if err != nil {
		// log.Fatal(err.Error())
		fmt.Printf("Error Occured\n%s\n", err.Error())
	} else {
		// if C:\\temp\\1.txt exists
		println(f.Name())
	}

	_, customError := requestError()

	if customError != nil {
		fmt.Printf("Error Occured\n%s\n", customError.Error())
	}
}
