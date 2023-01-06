package router

import (
	"api/internal/router/rotas"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	return rotas.ConfigRouter(r)
}
