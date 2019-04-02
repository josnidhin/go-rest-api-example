/**
 * @author Jose Nidhin
 */
package handlers

import (
	"net/http"

	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"

	"github.com/josnidhin/go-rest-api-example/system"
)

var app *system.App
var validate *validator.Validate

func New(inst *system.App) {
	app = inst
	validate = validator.New()
}

func Default404(w http.ResponseWriter, r *http.Request) {
	res := &apiError{}
	res.Status = http.StatusNotFound
	renderError(w, http.StatusNotFound, res)
}

func getConfig() *system.Config {
	return app.Config
}

func getLogger() *zap.Logger {
	return app.Logger
}
