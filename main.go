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

	a.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})
}

func (a *App) Get(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, fn).Methods(http.MethodGet)
}

func (a *App) Post(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, fn).Methods(http.MethodPost)
}

func (a *App) Start(config *config.Config) {
	http.ListenAndServe(fmt.Sprintf(":%d", config.Server.HTTP.Port), a.Router)
}
