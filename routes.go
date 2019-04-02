/**
 * @author Jose Nidhin
 */
package main

import (
	"github.com/josnidhin/go-rest-api-example/handlers"
	"github.com/josnidhin/go-rest-api-example/system"
)

var routes = system.Routes{
	system.Route{
		Method:      "GET",
		Path:        "/hello",
		HandlerFunc: handlers.Hello,
	},

	system.Route{
		Method:      "POST",
		Path:        "/hello",
		HandlerFunc: handlers.CustomHello,
	},
}

func Routes() system.Routes {
	return routes
}
