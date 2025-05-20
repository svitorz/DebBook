package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario models.Usuario

	if err = json.Unmarshal(body, &usuario); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewUsersRepository(db)

	usuarioSalvoNoBanco, err := repository.FindByEmail(usuario.Email)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(usuario.Senha, usuarioSalvoNoBanco.Senha); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		fmt.Println(err)
		return
	}

	token, err := auth.CreateToken(usuarioSalvoNoBanco.ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
