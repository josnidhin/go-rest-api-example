/**
 * @author Jose Nidhin
 */
package handlers

import (
	"github.com/josnidhin/go-rest-api-example/system"
	"go.uber.org/zap"
)

var app *system.App

func SetApp(inst *system.App) {
	app = inst
}

func getConfig() *system.Config {
	return app.Config
}

func getLogger() *zap.Logger {
	return app.Logger
}
