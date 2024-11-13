package main

func main() {
	for i := 0; i <= 10; i++ {
		println(i)
	}

	numbers := []string{"one", "two", "three"}
	for index, value := range numbers {
		println(index, value)
	}

	for _, value := range numbers {
		println(value)
	}

	for index := range numbers {
		println(index)
	}
}
