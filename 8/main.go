package main

import "fmt"

func main() {
	sumValue, isSumEven := sum(1, 3)
	fmt.Println(sumValue, isSumEven)
}

func sum(a int, b int) (int, bool) {
	sum := a + b
	isEven := sum%2 == 0
	return sum, isEven
}
