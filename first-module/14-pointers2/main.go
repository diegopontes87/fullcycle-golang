package main

func main() {

	a := 10
	//Accessing the address of a
	var pointer *int = &a
	println(a)
	//Changing the value inside the address
	*pointer = 20
	//Accessing the address of a through b
	b := &a
	println(a)
	println(*b)
	*b = 30
	println(*b)
	println(a)
	println(pointer)
}
