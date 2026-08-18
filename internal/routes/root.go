package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/repeale/fp-go"
)

// func fpGoUsecase() {
// 	// Create an Option
// 	some := option.Some(42)
// 	none := option.None[int]()

// 	// Map over values
// 	doubled := option.Map(N.Mul(2))(some)
// 	fmt.Println(option.GetOrElse(lazy.Of(0))(doubled)) // Output: 84
// 	fmt.Println(option.GetOrElse(lazy.Of(1))(none))    // Output: 1

// 	// // Chain operations
// 	result := option.Chain(func(x int) option.Option[string] {
// 		if x > 0 {
// 			return option.Some(fmt.Sprintf("Positive: %d", x))
// 		}
// 		return option.None[string]()
// 	})(some)
// 	fmt.Println(option.GetOrElse(lazy.Of("No value"))(result)) // Output: Positive: 42
// }

func RootPage(w http.ResponseWriter, r *http.Request) {
	// an example API handler

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

type SimpleDirEntry struct {
	FullPath string
	Name     string
	isDir    bool
}

func simplifyOsDirEntry(dir string, entry os.DirEntry) SimpleDirEntry {
	return SimpleDirEntry{
		FullPath: filepath.Join(dir, entry.Name()),
		Name:     entry.Name(),
		isDir:    entry.IsDir(),
	}
}

func readDirSimple(dir string) []SimpleDirEntry {
	osEntries, _ := os.ReadDir(dir)
	simplifyOsDirEntryCurrDir := fp.Curry2(simplifyOsDirEntry)(dir)
	return fp.Map(simplifyOsDirEntryCurrDir)(osEntries)
}

func walkingReadDir(dir string) []SimpleDirEntry {
	simpleEntries := readDirSimple(dir)
	return fp.Reduce(func(acc []SimpleDirEntry, curr SimpleDirEntry) []SimpleDirEntry {
		// Recursive read the dir with output SimpleDirEntry{}
		if curr.isDir {
			nextDirToScan := filepath.Join(dir, curr.Name)
			return append(acc, walkingReadDir(nextDirToScan)...)
		}

		// Return all the file list if there's no dir.
		return append(acc, curr)
	}, []SimpleDirEntry{})(simpleEntries)
}

// we only need call this once.
func ParseTemplateWithinDir(viewDir string) *template.Template {
	allFiles := walkingReadDir(viewDir)
	allTemplateFiles := fp.Map(func(a SimpleDirEntry) string {
		return a.FullPath
	})(allFiles)
	tmpl := template.Must(template.ParseFiles(allTemplateFiles...))
	return tmpl
}

// Ref: https://medium.com/@uygaroztcyln/clean-ui-with-gos-html-templates-base-partials-and-funcmaps-4915296c9097
func TemplatePage(w http.ResponseWriter, r *http.Request) {
	var viewDir, _ = filepath.Abs("templates/")
	var tmpl = ParseTemplateWithinDir(viewDir)
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
