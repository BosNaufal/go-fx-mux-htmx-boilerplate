package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", RootPage)

	return router
}
