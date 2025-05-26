package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetUserID(r)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	request, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publicacao models.Publicacao

	publicacao.AutorID = userId

	if err = json.Unmarshal(request, &publicacao); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	newPostRespository := repository.NewPostsRepository(db)
	publicacao.ID, err = newPostRespository.Store(publicacao)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publicacao)
}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request)    {}
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request)   {}
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {}
func DeletarPublicacao(w http.ResponseWriter, r *http.Request)   {}
