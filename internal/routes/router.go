package routes

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	// This will serve files under http://localhost:8000/static/<filename>
	publicDir, _ := filepath.Abs("public/")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(publicDir))))

	router.HandleFunc("/api", RootResource)
	router.HandleFunc("/", RootPage)
	return router
}
