package main

import "fmt"

func main() {

	payments := map[string]int{"Diego": 2000, "Renata": 9000}
	fmt.Println(payments["Diego"])
	delete(payments, "Diego")
	fmt.Println(payments["Diego"])
	payments["Diego"] = 5000
	fmt.Println(payments["Diego"])

	for name, payment := range payments {
		fmt.Printf("The payment for %s is %d\n", name, payment)
	}

	for _, payment := range payments {
		fmt.Printf("The payment is %d\n", payment)
	}

	newPayment := make(map[string]int)
	fmt.Println(newPayment)

}
