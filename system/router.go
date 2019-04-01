/**
 * @author Jose Nidhin
 */
package system

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/josnidhin/go-rest-api-example/config"
	"github.com/josnidhin/go-rest-api-example/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(requestLogger)

	router.NotFoundHandler = http.HandlerFunc(handlers.Default404)

	v1Router := router.PathPrefix("/v1").Subrouter()

	for _, route := range config.AppRoutes {
		v1Router.
			Methods(route.Method).
			Path(route.Path).
			HandlerFunc(route.HandlerFunc)
	}

	return router
}

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
