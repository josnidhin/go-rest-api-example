/**
 * @author Jose Nidhin
 */
package system

import (
	"go.uber.org/zap"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"

	"github.com/josnidhin/go-rest-api-example/system/config"
)

//
type App struct {
	Config *config.Config
	Logger *zap.Logger
	DB     *sqlx.DB
	Router *chi.Mux
}
