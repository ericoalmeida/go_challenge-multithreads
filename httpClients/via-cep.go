package httpClients

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getViaCEPURL(cep string) string {
	return "http://viacep.com.br/ws/" + cep + "/json/"
}

type CEPContent struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Ddd         string `json:"ddd"`
	Gia         string `json:"gia"`
	Siafi       string `json:"siafi"`
}

type ViaCEPData struct {
	Origin string
	Cep    CEPContent
}

var ViaCEPChannel = make(chan ViaCEPData)

func GetCEPInfoFromViaCEP(cep string) {
	response, err := http.Get(getViaCEPURL(cep))
	if err != nil {
		panic(err)
	}

	err = publishContent(response)
	if err != nil {
		panic(err)
	}
}

func publishContent(response *http.Response) error {
	cep := CEPContent{}
	data := ViaCEPData{}

	decoder := json.NewDecoder(response.Body)

	err := decoder.Decode(&cep)
	if err != nil {
		return err
	}

	data.Origin = fmt.Sprintf("%s [%s]", "Via CEP", "https://viacep.com.br/")
	data.Cep = cep

	ViaCEPChannel <- data

	return nil
}
