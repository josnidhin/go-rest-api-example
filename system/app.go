/**
 * @author Jose Nidhin
 */
package system

import (
	"database/sql"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type App struct {
	Config *Config
	Logger *zap.Logger
	DB     *sql.DB
	Router *mux.Router
}
