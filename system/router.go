/**
 * @author Jose Nidhin
 */
package system

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/josnidhin/go-rest-api-example/config"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(requestLogger)

	router.NotFoundHandler = http.HandlerFunc(config.Default404Handler)

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
		start := time.Now()
		next.ServeHTTP(w, r)

		AppLogger().Info(
			"Request log",
			zap.String("method", r.Method),
			zap.String("url", r.RequestURI),
			zap.String("userAgent", r.Header.Get("User-Agent")),
			zap.String("httpVersion", r.Proto),
			zap.Duration("requestDuration", time.Since(start)),
		)
	})
}
