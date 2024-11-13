package main

import "fmt"

type Address struct {
	City   string
	Number int
	State  string
}

type Person interface {
	Disable()
}

type Client struct {
	Name     string
	Age      int
	IsActive bool
	Address  Address
}

func (client Client) Disable() {
	client.IsActive = false
	fmt.Printf("The client %s was removed: %t\n", client.Name, !client.IsActive)
}

func main() {
	diego := Client{
		Name:     "Diego",
		Age:      36,
		IsActive: true,
	}

	diego.Address.City = "Belo Horizonte"

	fmt.Printf("Name: %s, age: %d, is active: %t\n", diego.Name, diego.Age, diego.IsActive)
	fmt.Printf("City: %s\n", diego.Address.City)

	diego.Disable()
}
