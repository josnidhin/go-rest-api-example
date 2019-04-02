/**
 * @author Jose Nidhin
 */
package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/josnidhin/go-rest-api-example/handlers"
	"github.com/josnidhin/go-rest-api-example/system"
)

func main() {
	app := &system.App{}

	app.Config = system.ConfigInstance()
	app.Logger = system.NewLogger(app.Config)

	handlers.New(app)
	app.Router = system.NewRouter(Routes(), handlers.Default404)

	serverAddress := fmt.Sprintf(":%d", app.Config.Server.HTTP.Port)
	err := http.ListenAndServe(serverAddress, app.Router)

	if err != nil {
		app.Logger.Fatal("Server startup failed", zap.Error(err))
	}
}
