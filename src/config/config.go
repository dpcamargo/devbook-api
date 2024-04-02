package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
	SecretKey          []byte
)

func Carregar() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=UTF8MB4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_USER_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_HOST_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	if SecretKey = []byte(os.Getenv("SECRET_KEY")); SecretKey == nil {
		log.Fatal("SECRET_KEY não encontrado em .env")
	}
}
