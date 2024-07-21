package main

func sum(a, b *int) int {
	*a = 25
	return *a + *b
}
func main() {
	var1 := 10
	var2 := 20

	println(sum(&var1, &var2))
	println(var1)
}
