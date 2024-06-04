package httpClients

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getBrasilAPIUrl(cep string) string {
	return "https://brasilapi.com.br/api/cep/v2/" + cep
}

type Coordinates struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type Location struct {
	Type        string      `json:"type"`
	Coordinates Coordinates `json:"coordinates"`
}

type CEPInfo struct {
	CEP          string   `json:"cep"`
	State        string   `json:"state"`
	City         string   `json:"City"`
	Neighborhood string   `json:"neighborhood"`
	Street       string   `json:"street"`
	Service      string   `json:"service"`
	Location     Location `json:"location"`
}

type BrasilAPIData struct {
	Origin string
	Cep    CEPInfo
}

var BrasilAPIChannel = make(chan BrasilAPIData)

func GetCEPInfoFromBrasilAPI(cep string) {
	response, err := http.Get(getBrasilAPIUrl(cep))
	if err != nil {
		panic(err)
	}

	err = publishInfo(response)
	if err != nil {
		panic(err)
	}
}

func publishInfo(response *http.Response) error {
	cep := CEPInfo{}
	data := BrasilAPIData{}

	decoder := json.NewDecoder(response.Body)

	err := decoder.Decode(&cep)
	if err != nil {
		return err
	}

	data.Origin = fmt.Sprintf("%s [%s]", "Brasil API", "https://brasilapi.com.br/")
	data.Cep = cep

	BrasilAPIChannel <- data

	return nil
}
