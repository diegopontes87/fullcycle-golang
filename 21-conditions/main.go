package main

func main() {
	a := 23
	b := 2
	c := 4

	if a > b {
		println(a)
	} else {
		println(b)
	}

	if a > b && c > a {
		println("a > b && c > a")
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	default:
		println("default")
	}
}
