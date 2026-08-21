package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/BosNaufal/go-fx-mux-htmx-boilerplate/internal/helpers"
	"github.com/BosNaufal/go-fx-mux-htmx-boilerplate/internal/views"
)

type SearchResultItem struct {
	Title       string
	Link        string
	Description string
}

type RootPageData struct {
	SomeData    string
	SearchQuery string
	Results     []SearchResultItem
}

func generateDummyArticleTitle() string {
	return fmt.Sprintf("Some title Article %d", time.Now().UnixNano())
}

// Ref: https://medium.com/@uygaroztcyln/clean-ui-with-gos-html-templates-base-partials-and-funcmaps-4915296c9097
func RootPage(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	searchQuery := queries.Get("q")

	err := views.RenderPage(w, "home.html", RootPageData{
		SomeData:    "someData to render",
		SearchQuery: searchQuery,
		Results: []SearchResultItem{
			{
				Title:       generateDummyArticleTitle(),
				Link:        "https://to-real.com/link/article",
				Description: "Some long and very long article descriptions that we won't to read",
			},
			{
				Title:       generateDummyArticleTitle(),
				Link:        "https://to-real.com/link/article",
				Description: "Some long and very long article descriptions that we won't to read",
			},
			{
				Title:       generateDummyArticleTitle(),
				Link:        "https://to-real.com/link/article",
				Description: "Some long and very long article descriptions that we won't to read",
			},
		},
	})

	if err != nil {
		helpers.ResponseString(w, "something went wrong")
		return
	}
}
