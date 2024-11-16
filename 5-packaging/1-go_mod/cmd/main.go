package main

import (
	"fmt"
	"fullcycle_golang/5-packaging/1-go_mod/math"
)

func main() {
	res := math.Math{NumberA: 1, NumberB: 3}.Add()
	fmt.Println(res)
}
