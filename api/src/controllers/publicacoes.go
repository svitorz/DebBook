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
	"strconv"

	"github.com/gorilla/mux"
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
		responses.Err(w, http.StatusForbidden, err)
		return
	}

	if err := publicacao.Preparar(); err != nil {
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

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusNoContent, err)
		return
	}
	db, err := database.Conectar()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewPostsRepository(db)
	post, err := repository.FindById(postId)
	if err != nil {
		responses.Err(w, http.StatusNotFound, err)
		return
	}

	responses.JSON(w, http.StatusOK, post)
}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()
	repository := repository.NewPostsRepository(db)
	posts, err := repository.Index(userId)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetUserID(r)
	if err != nil {
		responses.Err(w, http.StatusForbidden, err)
		return
	}

	params := mux.Vars(r)
	postId, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusNotFound, err)
	}

	db, err := database.Conectar()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewPostsRepository(db)
	savedPost, err := repository.FindById(postId)
	if err != nil {
		responses.Err(w, http.StatusNotFound, err)
	}

	if savedPost.AutorID != userId {
		responses.Err(w, http.StatusForbidden, err)
		return
	}

	request, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	var post models.Publicacao
	if err = json.Unmarshal(request, &post); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = post.Preparar(); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if err = repository.Update(postId, post); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetUserID(r)
	if err != nil {
		responses.Err(w, http.StatusForbidden, err)
		return
	}

	params := mux.Vars(r)
	postId, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusNotFound, err)
	}

	db, err := database.Conectar()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewPostsRepository(db)
	savedPost, err := repository.FindById(postId)
	if err != nil {
		responses.Err(w, http.StatusNotFound, err)
	}

	if savedPost.AutorID != userId {
		responses.Err(w, http.StatusForbidden, err)
		return
	}

	request, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	var post models.Publicacao
	if err = json.Unmarshal(request, &post); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = post.Preparar(); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if err = repository.Destroy(postId); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func BuscarPublicacaoDoUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["usuarioId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusNoContent, err)
		return
	}
	db, err := database.Conectar()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewPostsRepository(db)
	post, err := repository.FindByUserId(postId)
	if err != nil {
		responses.Err(w, http.StatusNotFound, err)
		return
	}

	responses.JSON(w, http.StatusOK, post)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := auth.GetUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	postId, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusNoContent, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()
	repository := repository.NewPostsRepository(db)

	savedPost, err := repository.FindById(postId)
	if err != nil {
		responses.Err(w, http.StatusNotFound, err)
	}

	if err = repository.LikePost(userId, savedPost.ID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := auth.GetUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	postId, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusNoContent, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()
	repository := repository.NewPostsRepository(db)

	savedPost, err := repository.FindById(postId)
	if err != nil {
		responses.Err(w, http.StatusNotFound, err)
	}

	if err = repository.UnlikePost(userId, savedPost.ID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
