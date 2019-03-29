/**
 * @author Jose Nidhin
 */
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/josnidhin/go-rest-api-example/config"
	"github.com/josnidhin/go-rest-api-example/libs"
	"github.com/josnidhin/go-rest-api-example/routes"
)

type App struct {
	Config *config.Config
	Logger *zap.Logger
	Router *mux.Router
}

func main() {
	app := &App{}
	app.Config = config.Load()

	app.Initilise()
	app.Start()
}

func (a *App) Initilise() {
	a.Logger = libs.NewLogger(a.Config)
	a.Router = routes.NewRouter()
}

func (a *App) Start() {
	serverAddress := fmt.Sprintf(":%d", a.Config.Server.HTTP.Port)

	err := http.ListenAndServe(serverAddress, a.Router)

	if err != nil {
		a.Logger.Fatal("Server startup failed", zap.Error(err))
	}
}
