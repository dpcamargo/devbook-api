package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var usuario modelos.Usuario
	if err = json.Unmarshal(corpoRequest, &usuario); err != nil {
		log.Fatal(err)
	}

	db, err := banco.Conectar()
	if err != nil {
		log.Fatal(err)
	}
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioID, err := repositorio.Criar(usuario)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("ID inserido: %d", usuarioID)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca todos usuários"))
}
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário"))
}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
