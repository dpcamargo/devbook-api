package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/router"
)

func main() {
	fmt.Println("Iniciando o servidor...")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}
