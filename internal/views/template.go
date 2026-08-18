package views

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/BosNaufal/go-fx-mux-htmx-boilerplate/internal/helpers"
	"github.com/repeale/fp-go"
)

func parseTemplateFiles(filepaths []string) *template.Template {
	tmpl := template.Must(template.ParseFiles(filepaths...))
	return tmpl
}

func renderPageHTML(pageFileName string) *template.Template {
	var layoutDir, _ = filepath.Abs("templates/layouts")

	layoutFiles := fp.Pipe2(
		helpers.WalkingReadDir,
		fp.Map(helpers.SimpleDirEntryFullPath),
	)(layoutDir)

	var pageDir, _ = filepath.Abs("templates/pages")
	var pageTargetPath = filepath.Join(pageDir, pageFileName)
	var tmpl = parseTemplateFiles(append(layoutFiles, pageTargetPath))
	return tmpl
}

// When need custom layout
func RenderPageWithLayout(writer io.Writer, layoutName string, absolutePathToPageHTML string, data any) error {
	var tmpl = renderPageHTML(absolutePathToPageHTML)
	err := tmpl.ExecuteTemplate(writer, layoutName, data)
	return err
}

var defaultLayout = "base_layout.html"

// Using default layout
func RenderPage(writer io.Writer, absolutePathToPageHTML string, data any) error {
	var tmpl = renderPageHTML(absolutePathToPageHTML)
	err := tmpl.ExecuteTemplate(writer, defaultLayout, data)
	return err
}
