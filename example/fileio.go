package example

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ExampleFile() {
	// open file with READ/WRITE 0600 permission
	fileOpen, errOpen := os.OpenFile(".\\text.txt", os.O_RDWR, 0600)
	if errOpen != nil {
		// create file if not exists
		fileCreate, errCreate := os.Create(".\\text.txt")
		// when create failed
		if errCreate != nil {
			// panic
			panic(errCreate)
		}
		// assign pointer to fileOpen
		fileOpen = fileCreate
		fileCreate = nil
	}
	// defer close fileOpen
	defer fileOpen.Close()

	// 64 byte read buffer
	buffer := make([]byte, 64)
	// concatenated text buffer
	bufferText := []byte{}
	// readed bytes length
	readedBytes := 0

	for {
		// read 64 bytes buffer from file
		count, errRead := fileOpen.Read(buffer)
		// error occred and not End of File
		if errRead != nil && errRead != io.EOF {
			// panic
			panic(errRead)
		}

		// No more readable buffer
		if count == 0 {
			break
		}

		// add readed bytes length
		readedBytes += count
		// concat byte buffer
		bufferText = append(bufferText, buffer[:count]...)
	}
	// assert []byte to string
	readedText := string(bufferText)
	// print readed text
	fmt.Printf("readed text : %s\n", readedText)

	// write text
	textContent := "Hello World! "
	// if txt is not empty
	if len(readedText) != 0 {
		// parse readed text and split
		numText := strings.Split(readedText, " ")[2]
		// "Hello World! {num}" << read and parse num
		num, errConvert := strconv.ParseInt(numText, 0, 64)
		// if num not exists or other error occured
		if errConvert != nil {
			num = 1
		}
		// concat string and int
		textContent = fmt.Sprint(textContent, num+1)
	}
	// assert string to []byte
	textBytes := []byte(textContent)

	// write text
	num, errWrite := fileOpen.WriteAt(textBytes, 0)
	// error occured
	if errWrite != nil {
		// panic
		panic(errWrite)
	}
	fmt.Printf("write %d bytes\n", num)
}
