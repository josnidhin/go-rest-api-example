/**
 * @author Jose Nidhin
 */
package system

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type App struct {
	Config *Config
	Logger *zap.Logger
	Router *mux.Router
}
