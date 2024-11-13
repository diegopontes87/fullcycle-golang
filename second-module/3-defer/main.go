package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	//It's going to be the last thing to be called after the whole execution
	defer request.Body.Close()

	result, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	println(string(result))
	request.Body.Close()
}
