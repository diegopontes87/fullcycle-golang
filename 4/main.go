package main

import "fmt"

type ID int

var (
	e ID = 1
)

func main() {
	fmt.Printf("o tipo de da variavel é: %T\n", e)
}
