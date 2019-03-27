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
	app.Init(config)
}

func (app *App)Init(config *config.Config) {
	app.Router = mux.NewRouter()

	app.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})
}

func (app *App)Get(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, fn).Methods(http.MethodGet)
}

func (app *App)Post(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, fn).Methods(http.MethodPost)
}
