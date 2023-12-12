package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Info struct {
	Nome string `json:"nome"`
}

func main() {
	addMedicamento()
}

func initialization() int {
	var option int
	fmt.Println("aqui você pode adicionar medicamentos ao estoque")
	fmt.Println("para continuar tecle 1 para sair tecle 0")
	fmt.Scanln(&option)
	return option
}

func addMedicamento() {
	var medicamento string
	switch initialization() {
	case 1:
		fmt.Println("digite o nome do medicamente a ser adicionado")
		fmt.Scanln(&medicamento)
	case 0:
		fmt.Println("saindo do programa")
		os.Exit(0)
	}

	info := Info{Nome: medicamento}

	jsonData, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("remedios.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Informações adicionadas ao arquivo JSON com sucesso!")

}
