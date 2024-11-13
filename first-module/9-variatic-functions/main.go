package main

import "fmt"

func main() {
	fmt.Println(sum(2, 3, 4, 52, 3, 51))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
