package example

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func requestGetMethod() {
	// get request to "https://jsonplaceholder.typicode.com/todos/1"

	request, errRequest := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos/1", nil)
	if errRequest != nil {
		panic(errRequest)
	}

	request.Header.Add("User-Agent", "Chrome")
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	response, errGet := client.Do(request)

	// simple way
	// response, errGet := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if errGet != nil {
		panic(errGet)
	}
	// defer close body stream
	defer response.Body.Close()

	// 1024 readable buffer
	buffer := make([]byte, 1024)
	bufferConcatenated := []byte{}
	for {
		// read buffer from response body stream
		count, errRead := response.Body.Read(buffer)

		if errRead != nil && errRead != io.EOF {
			panic(errRead)
		}

		if count == 0 {
			break
		}

		// concat buffer
		bufferConcatenated = append(bufferConcatenated, buffer[:count]...)
	}
	// assert concatenated text and print
	bufferText := string(bufferConcatenated)
	fmt.Printf("%s\n", bufferText)
}

func requestPost() {
	bodyContent := bytes.NewBufferString("Hello World")
	response, errResponse := http.Post("https://httpbin.org/post", "text/plain", bodyContent)
	if errResponse != nil {
		panic(errResponse)
	}

	defer response.Body.Close()

	responseBody, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		fmt.Printf("Error occured\n%s", errRead.Error())
	}

	fmt.Println("response : ", string(responseBody))
}

func requestPostForm() {
	bodyContent := url.Values{
		"Name": {"Hong Gildong"},
		"Age":  {"26"},
	}
	response, errResponse := http.PostForm("https://httpbin.org/post", bodyContent)
	if errResponse != nil {
		panic(errResponse)
	}

	defer response.Body.Close()

	responseBody, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		fmt.Printf("Error occured\n%s", errRead.Error())
	}

	fmt.Println("response : ", string(responseBody))
}

func requestPostMethod() {
	requestPost()
	requestPostForm()
}

func ExampleRequest() {
	requestGetMethod()
	requestPostMethod()
}
