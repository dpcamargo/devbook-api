package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

func main() {
	config.Carregar()
	fmt.Printf("Running on http://localhost:%d\n", config.Porta)
	fmt.Println("String de conexão com o banco de dados:", config.StringConexaoBanco)
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
