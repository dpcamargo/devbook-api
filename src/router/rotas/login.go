package rotas

import (
	"net/http"

	"api/src/controllers"
)

var RotaLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
