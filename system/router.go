/**
 * @author Jose Nidhin
 */
package system

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(routes Routes, notFoundHandler http.HandlerFunc) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(RequestLogger)

	router.NotFoundHandler = notFoundHandler

	v1Router := router.PathPrefix("/v1").Subrouter()

	for _, route := range routes {
		v1Router.
			Methods(route.Method).
			Path(route.Path).
			HandlerFunc(route.HandlerFunc)
	}

	return router
}
