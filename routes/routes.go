/**
 * @author Jose Nidhin
 */
package routes

import (
	"net/http"

	"github.com/josnidhin/go-rest-api-example/handlers"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Method:      "GET",
		Path:        "/hello",
		HandlerFunc: handlers.Hello,
	},

	Route{
		Method:      "POST",
		Path:        "/hello",
		HandlerFunc: handlers.CustomHello,
	},
}
