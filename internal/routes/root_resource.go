package routes

import (
	"net/http"

	"github.com/BosNaufal/go-fx-mux-htmx-boilerplate/internal/helpers"
)

func RootResource(w http.ResponseWriter, r *http.Request) {
	// an example API handler
	helpers.ResponseJSON(w, map[string]bool{"ok": true})
}
