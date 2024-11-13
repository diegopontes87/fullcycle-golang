package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func main() {

	for _, cep := range os.Args[1:] {
		url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error during request: %v\n", err)
			continue
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading response%v\n", err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error during parse%v\n", err)
		}
		fmt.Println(data)

		file, err := os.Create("city.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file%v\n", err)
		}

		defer file.Close()
		file.WriteString(fmt.Sprintf("CEP: %s, Localidade:%s", data.CEP, data.Localidade))
	}

}
