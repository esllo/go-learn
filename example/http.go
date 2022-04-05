package example

import (
	"io/ioutil"
	"net/http"
)

func handleFuncExample() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world"))
	})
}

func handleExample() {
	http.Handle("/path", new(helloHandler))
}

type helloHandler struct{}

func (handler *helloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(request.URL.Path))
}

type staticHandler struct{}

func (handler *staticHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	filePath := "text.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		writer.WriteHeader(404)
		writer.Write([]byte(http.StatusText(404)))
		return
	}

	writer.Header().Add("Content-Type", "text/plain")
	writer.Write(content)
}

func handleFileExample() {
	http.Handle("/text", new(staticHandler))
}

func ExampleHttp() {
	handleFuncExample()
	handleExample()
	handleFileExample()
	http.ListenAndServe(":5000", nil)
}
