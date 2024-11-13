package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number int `json:"number"`
	Amount int `json:"amount"`
}

func main() {

	account := Account{Number: 1, Amount: 100}

	result, err := json.Marshal(account)

	if err != nil {
		println(err)
	}
	println(string(result))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		println(err)
	}

	pureJson := []byte(`{"number": 1, "amount": 200}`)

	var accountX Account
	err = json.Unmarshal(pureJson, &accountX)
	if err != nil {
		println(err)
	}
	println(accountX.Number)
	println(accountX.Amount)

}
