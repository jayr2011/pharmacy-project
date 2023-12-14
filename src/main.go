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
	Mg   string `json:"mg"`
}

func main() {
	addMedicamento()
}

func initialization() int {
	var option int
	fmt.Println("aqui você pode adicionar medicamentos ao estoque e consultar os medicamentos disponiveis")
	fmt.Println("tecle 1 para adicionar")
	fmt.Println("tecle 2 para consultar os medicamentos do estoque")
	fmt.Println("tecle 0 para sair do programa")
	fmt.Scanln(&option)
	return option
}

func addMedicamento() {
	var medicamento string
	var mg string
	switch initialization() {
	case 1:
		fmt.Println("digite o nome do medicamente a ser adicionado")
		fmt.Scanln(&medicamento)
		fmt.Println("digite a quantidade de mg do medicamento")
		fmt.Scanln(&mg)
	case 2:
		fmt.Println("consultando medicamentos")
		leMedicamentos()
	case 0:
		fmt.Println("saindo do programa")
		os.Exit(0)
	}

	info := Info{Nome: medicamento, Mg: mg}

	dadosDoJson, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("remedios.json", dadosDoJson, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Informações adicionadas ao arquivo JSON com sucesso!")

}

func leMedicamentos() {
	file, err := ioutil.ReadFile("remedios.json")
	if err != nil {
		log.Fatal(err)
	}

	var info Info

	err = json.Unmarshal(file, &info)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Nome: ", info.Nome)
	fmt.Println("Mg: ", info.Mg)
}
