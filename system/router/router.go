/**
 * @author Jose Nidhin
 */
package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

//
type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

//
type Routes []Route

//
func New(routes Routes, notFoundHandler http.HandlerFunc) *chi.Mux {
	router := chi.NewRouter()

	//router.Use(RequestLogger)

	router.NotFound(notFoundHandler)

	v1Router := chi.NewRouter()

	for _, route := range routes {
		v1Router.
			MethodFunc(
				route.Method,
				route.Path,
				route.HandlerFunc)
	}

	router.Mount("/v1", v1Router)

	return router
}
