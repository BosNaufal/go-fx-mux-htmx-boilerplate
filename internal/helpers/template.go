package helpers

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/repeale/fp-go"
)

func parseTemplateFiles(filepaths []string) *template.Template {
	tmpl := template.Must(template.ParseFiles(filepaths...))
	return tmpl
}

func renderPageHTML(pageFileName string) *template.Template {
	var layoutDir, _ = filepath.Abs("templates/layouts")

	layoutFiles := fp.Pipe2(
		WalkingReadDir,
		fp.Map(SimpleDirEntryFullPath),
	)(layoutDir)

	var pageDir, _ = filepath.Abs("templates/pages")
	var pageTargetPath = filepath.Join(pageDir, pageFileName)
	var tmpl = parseTemplateFiles(append(layoutFiles, pageTargetPath))
	return tmpl
}

func RenderPageWithLayout(writer io.Writer, layoutName string, absolutePathToPageHTML string, data any) error {
	var tmpl = renderPageHTML(absolutePathToPageHTML)
	err := tmpl.ExecuteTemplate(writer, layoutName, data)
	return err
}
