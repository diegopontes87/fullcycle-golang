package main

import "fmt"

func main() {

	var myVar interface{} = "Hello World"
	println(myVar.(string))

	result, isParsed := myVar.(string)
	if isParsed {
		fmt.Printf("The value of result is: %v and isParsed: %v\n", result, isParsed)
	}

	res, isParsed := myVar.(int)
	fmt.Printf("The value of result is: %v and isParsed: %v\n", res, isParsed)

}
