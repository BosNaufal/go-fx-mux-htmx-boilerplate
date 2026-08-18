package routes

import (
	"net/http"

	"github.com/BosNaufal/go-fx-mux-htmx-boilerplate/internal/helpers"
	"github.com/BosNaufal/go-fx-mux-htmx-boilerplate/internal/views"
)

type RootPageData struct {
	SomeData string
}

// Ref: https://medium.com/@uygaroztcyln/clean-ui-with-gos-html-templates-base-partials-and-funcmaps-4915296c9097
func RootPage(w http.ResponseWriter, r *http.Request) {
	err := views.RenderPage(w, "home.html", RootPageData{
		SomeData: "someData to render",
	})

	if err != nil {
		helpers.ResponseString(w, "something went wrong")
		return
	}
}
