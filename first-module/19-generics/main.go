package main

func SumInteger(newMap map[string]int) int {
	var sum int
	for _, value := range newMap {
		sum += value
	}
	return sum
}

func SumFloat(newMap map[string]float64) float64 {
	var sum float64
	for _, value := range newMap {
		sum += value
	}
	return sum
}

type NewNumber int

type Number interface {
	~int | ~float64
}

func Sum[T Number](newMap map[string]T) T {
	var sum T
	for _, value := range newMap {
		sum += value
	}
	return sum
}

func CompareNumbers[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Diego": 1000, "Renata": 2000}
	m2 := map[string]float64{"Diego": 100.0, "Renata": 200.0}
	m3 := map[string]NewNumber{"Diego": 100.0, "Renata": 200.0}
	println(SumInteger(m))
	println(SumFloat(m2))
	println(Sum(m))
	println(Sum(m2))
	println(Sum(m3))
	println(CompareNumbers(10, 10.0))
}
