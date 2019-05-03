/**
 * @author Jose Nidhin
 */
package system

import (
	"database/sql"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type App struct {
	Config *Config
	Logger *zap.Logger
	DB     *sql.DB
	Router *chi.Mux
}
