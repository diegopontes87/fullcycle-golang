package main

import "fmt"

type Client struct {
	name string
}

func (c *Client) changeName() {
	c.name = "Diego Borges Pontes"
	fmt.Printf("Client new name is %v\n", c.name)
}

func main() {
	diego := Client{
		name: "Diego",
	}
	diego.changeName()
	fmt.Printf("Client name is %v\n", diego.name)

}
