/**
 * @author Jose Nidhin
 */
package handlers

import (
	"net/http"

	"go.uber.org/zap"
	"github.com/go-playground/validator/v10"

	"github.com/josnidhin/go-rest-api-example/system"
	"github.com/josnidhin/go-rest-api-example/system/config"
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

func getConfig() *config.Config {
	return app.Config
}

func getLogger() *zap.Logger {
	return app.Logger
}
