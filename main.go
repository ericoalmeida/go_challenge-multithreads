package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/ericoalmeida/go_challenges-multithreads/httpClients"
)

func main() {
	cep := flag.String("cep", "78590000", "CEP to get information about")
	flag.Parse()

	go httpClients.GetCEPInfoFromViaCEP(*cep)
	go httpClients.GetCEPInfoFromBrasilAPI(*cep)

	select {
	case cepData := <-httpClients.ViaCEPChannel:
		cepContent, _ := json.Marshal(cepData.Cep)

		fmt.Printf("Received content from.: %s\n\n", cepData.Origin)
		fmt.Printf("Content.: %s\n\n", cepContent)
	case cepData := <-httpClients.BrasilAPIChannel:
		cepContent, _ := json.Marshal(cepData.Cep)

		fmt.Printf("Received content from.: %s\n\n", cepData.Origin)
		fmt.Printf("Content.: %s\n\n", cepContent)
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}
