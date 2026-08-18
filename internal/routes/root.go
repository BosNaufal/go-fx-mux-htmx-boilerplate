package routes

import (
	"net/http"

	"github.com/BosNaufal/go-fx-mux-htmx-boilerplate/internal/helpers"
	"github.com/gorilla/mux"
)

func RootPage(w http.ResponseWriter, r *http.Request) {
	// an example API handler
	helpers.ResponseJSON(w, map[string]bool{"ok": true})
}

// Ref: https://medium.com/@uygaroztcyln/clean-ui-with-gos-html-templates-base-partials-and-funcmaps-4915296c9097
func TemplatePage(w http.ResponseWriter, r *http.Request) {
	err := helpers.RenderPage(w, "home.html", map[string]interface{}{
		"someData": "someData to render",
	})

	if err != nil {
		helpers.ResponseString(w, "something went wrong")
		return
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", RootPage)
	router.HandleFunc("/", TemplatePage)
	return router
}
