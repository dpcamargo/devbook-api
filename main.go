package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

// Geração de Secret com crypto.rand e base64 encoding
// func init() {
// 	chave := make([]byte, 64)

// 	if _, err := rand.Read(chave); err != nil {
// 		log.Fatal(err)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Println(stringBase64)
// }

func main() {
	config.Carregar()
	r := router.Gerar()

	fmt.Println(config.StringConexaoBanco)
	fmt.Printf("Running on http://localhost:%d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
