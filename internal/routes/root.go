package routes

import (
	"encoding/json"
	"net/http"

	"github.com/BosNaufal/go-fx-mux-htmx-boilerplate/internal/helpers"
	"github.com/gorilla/mux"
)

func RootPage(w http.ResponseWriter, r *http.Request) {
	// an example API handler
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

// Ref: https://medium.com/@uygaroztcyln/clean-ui-with-gos-html-templates-base-partials-and-funcmaps-4915296c9097
func TemplatePage(w http.ResponseWriter, r *http.Request) {
	err := helpers.RenderPageWithLayout(w, "base_layout.html", "templates/pages/home.html", map[string]interface{}{
		"someData": "someData to render",
	})

	if err != nil {
		w.Write([]byte("error nih bous"))
		return
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", RootPage)
	router.HandleFunc("/", TemplatePage)
	return router
}
