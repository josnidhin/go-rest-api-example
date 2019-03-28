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
	"github.com/josnidhin/go-rest-api-example/handlers"
)

type App struct {
	Router *mux.Router
}

func main() {
	config := config.Load()

	configJson, _ := json.Marshal(config)
	fmt.Println(string(configJson))

	app := &App{}
	app.Initilise(config)
	app.Start(config)
}

func (a *App) Initilise(config *config.Config) {
	a.Router = mux.NewRouter()

	a.Router.NotFoundHandler = http.HandlerFunc(handlers.Default404)
	a.Get("/", handlers.SayHello)
}

func (a *App) Get(path string, fn http.HandlerFunc) {
	a.Router.HandleFunc(path, fn).Methods(http.MethodGet)
}

func (a *App) Post(path string, fn http.HandlerFunc) {
	a.Router.HandleFunc(path, fn).Methods(http.MethodPost)
}

func (a *App) Start(config *config.Config) {
	http.ListenAndServe(fmt.Sprintf(":%d", config.Server.HTTP.Port), a.Router)
}
