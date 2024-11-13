package main

import "fmt"

func main() {

	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println("Array:", myArray)
	fmt.Println(len(myArray))
	for i, v := range myArray {
		fmt.Printf("index is %d and value is %d\n", i, v)
	}
}
