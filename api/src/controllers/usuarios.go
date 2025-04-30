package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usuario models.Usuario
	if err = json.Unmarshal(request, &usuario); err != nil {
		log.Fatal(err)
	}

	db, err := database.Conectar()
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewUsersRepository(db)
	userId, err := repository.Store(usuario)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", userId)))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuários"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualiando usuário"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
