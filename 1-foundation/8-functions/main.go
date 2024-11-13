package main

import (
	"errors"
	"fmt"
)

func main() {
	sumValue, isSumEven := sum(1, 3)
	fmt.Println(sumValue, isSumEven)

	result, error := divide(2, 0)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(result)
}

func sum(a int, b int) (int, bool) {
	sum := a + b
	isEven := sum%2 == 0
	return sum, isEven
}

func divide(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("Number cant be zero")
	}
	return a / b, nil
}
