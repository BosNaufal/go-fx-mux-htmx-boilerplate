package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", RootResource)
	router.HandleFunc("/", RootPage)
	return router
}
