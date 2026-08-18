package helpers

import (
	"text/template"

	"github.com/repeale/fp-go"
)

// we only need call this once.
func ParseTemplateWithinDir(viewDir string) *template.Template {
	allFiles := WalkingReadDir(viewDir)
	allTemplateFiles := fp.Map(func(a SimpleDirEntry) string {
		return a.FullPath
	})(allFiles)
	tmpl := template.Must(template.ParseFiles(allTemplateFiles...))
	return tmpl
}
