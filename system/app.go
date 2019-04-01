/**
 * @author Jose Nidhin
 */
package system

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/josnidhin/go-rest-api-example/config"
)

var app *App
var once sync.Once

type App struct {
	Config *config.Config
	Logger *zap.Logger
	Router *mux.Router
}

func (a *App) Initilise() {
	a.Logger = NewLogger(a.Config)
	a.Router = NewRouter()
}

func (a *App) Start() {
	serverAddress := fmt.Sprintf(":%d", a.Config.Server.HTTP.Port)

	err := http.ListenAndServe(serverAddress, a.Router)

	if err != nil {
		a.Logger.Fatal("Server startup failed", zap.Error(err))
	}
}

func AppInstance() *App {
	once.Do(func() {
		app = &App{}
		app.Config = config.Instance()

		app.Initilise()
	})

	return app
}

func AppLogger() *zap.Logger {
	return app.Logger
}
