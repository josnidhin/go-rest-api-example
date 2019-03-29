/**
 * @author Jose Nidhin
 */
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josnidhin/go-rest-api-example/config"
	"github.com/josnidhin/go-rest-api-example/routes"
)

type App struct {
	Config *config.Config
	Router *mux.Router
}

func main() {
	app := &App{}
	app.Config = config.Load()

	configJson, _ := json.Marshal(app.Config)
	fmt.Println(string(configJson))

	app.Initilise()
	app.Start()
}

func (a *App) Initilise() {
	a.Router = routes.NewRouter()
}

func (a *App) Start() {
	http.ListenAndServe(fmt.Sprintf(":%d", a.Config.Server.HTTP.Port), a.Router)
}
