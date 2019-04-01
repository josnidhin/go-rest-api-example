/**
 * @author Jose Nidhin
 */
package main

import (
	"github.com/josnidhin/go-rest-api-example/system"
)

func main() {
	app := system.AppInstance()
	app.Start()
}
