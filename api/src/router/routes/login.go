package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoutes = Route{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
