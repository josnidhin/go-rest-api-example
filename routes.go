/**
 * @author Jose Nidhin
 */
package main

import (
	"github.com/josnidhin/go-rest-api-example/handlers"
	"github.com/josnidhin/go-rest-api-example/system/router"
)

//
var routes = router.Routes{
	router.Route{
		Method:      "GET",
		Path:        "/hello",
		HandlerFunc: handlers.Hello,
	},

	router.Route{
		Method:      "POST",
		Path:        "/hello",
		HandlerFunc: handlers.CustomHello,
	},
}

//
func Routes() router.Routes {
	return routes
}
