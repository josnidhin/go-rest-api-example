/**
 * @author Jose Nidhin
 */
package system

import (
	"database/sql"

	"go.uber.org/zap"

	"github.com/go-chi/chi"

	"github.com/josnidhin/go-rest-api-example/system/config"
)

//
type App struct {
	Config *config.Config
	Logger *zap.Logger
	DB     *sql.DB
	Router *chi.Mux
}
