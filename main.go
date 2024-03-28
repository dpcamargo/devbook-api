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
	r := router.Gerar()
	
	fmt.Printf("Running on http://localhost:%d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
