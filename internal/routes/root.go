package routes

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"gitlab.com/bosnaufal/bos-ai-search/internal/helpers"
)

func RootPage(w http.ResponseWriter, r *http.Request) {
	// an example API handler
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

// Ref: https://medium.com/@uygaroztcyln/clean-ui-with-gos-html-templates-base-partials-and-funcmaps-4915296c9097
func TemplatePage(w http.ResponseWriter, r *http.Request) {
	var viewDir, _ = filepath.Abs("templates/")
	var tmpl = helpers.ParseTemplateWithinDir(viewDir)
	err := tmpl.ExecuteTemplate(w, "home.html", map[string]interface{}{
		"someData": "someData to render",
	})

	if err != nil {
		w.Write([]byte("error nih bous"))
		return
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/", RootPage)
	router.HandleFunc("/", TemplatePage)
	return router
}
