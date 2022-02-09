package main

import (
	"fmt"

	"github.com/arthuravila26/go-rest-api/database"
	"github.com/arthuravila26/go-rest-api/models"
	"github.com/arthuravila26/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
