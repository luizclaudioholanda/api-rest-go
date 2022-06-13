package main

import (
	"fmt"

	"github.com/luizclaudioholanda/go-rest-api/database"
	"github.com/luizclaudioholanda/go-rest-api/models"
	"github.com/luizclaudioholanda/go-rest-api/routes"
)

func main() {

	database.ConectaBancoDeDados()

	fmt.Println("Iniciando o servidor REST com GO")

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	routes.HandleRequest()
}
