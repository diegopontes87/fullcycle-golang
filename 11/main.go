package main

import "fmt"

type Client struct {
	Name     string
	Age      int
	IsActive bool
}

func main() {
	diego := Client{
		Name:     "Diego",
		Age:      36,
		IsActive: true,
	}

	fmt.Printf("Name: %s, age: %d, is active: %t\n", diego.Name, diego.Age, diego.IsActive)
}
