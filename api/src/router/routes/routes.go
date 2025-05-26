package routes

import (
	"api/src/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoutes)
	routes = append(routes, postRoutes...)
	for _, route := range routes {

		if route.RequerAutenticacao {
			r.HandleFunc(route.URI,
				middleware.Logger(
					middleware.Autenticar(route.Funcao)),
			).Methods(route.Metodo)
		} else {
			r.HandleFunc(route.URI, route.Funcao).Methods(route.Metodo)
		}

		r.HandleFunc(route.URI, route.Funcao).Methods(route.Metodo)
	}

	return r
}
