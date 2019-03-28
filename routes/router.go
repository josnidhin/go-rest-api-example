/**
 * @author Jose Nidhin
 */
package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josnidhin/go-rest-api-example/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.NotFoundHandler = http.HandlerFunc(handlers.Default404)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			HandlerFunc(route.HandlerFunc)
	}

	return router
}
