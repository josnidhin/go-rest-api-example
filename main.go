/**
 * @author Jose Nidhin
 */
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	//"github.com/josnidhin/go-rest-api-example/handlers"
	"github.com/josnidhin/go-rest-api-example/system"
	"github.com/josnidhin/go-rest-api-example/system/config"
	"github.com/josnidhin/go-rest-api-example/system/logger"
	"github.com/josnidhin/go-rest-api-example/system/pgdb"
	//"github.com/josnidhin/go-rest-api-example/system/router"
)

const (
	ShutdownTimeout = 10 * time.Second
)

var AppName, AppVersion string

func main() {
	app := &system.App{}
	server := &http.Server{}
	idleConnClosed := make(chan struct{})

	app.Config = config.New()
	app.Config.Log.AppName = AppName
	app.Config.Log.AppVersion = AppVersion

	app.Logger = logger.New(app.Config.Log)

	go signalHandler(server, idleConnClosed, app.Logger)

	app.DB = pgdb.New(app.Config.PG)

	//handlers.New(app)
	//app.Router = system.NewRouter(Routes(), handlers.Default404)

	serverAddress := fmt.Sprintf(":%d", app.Config.Server.HTTP.Port)

	server.Addr = serverAddress
	//server.Handler = app.Router

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		app.Logger.Fatal("Server startup failed", zap.Error(err))
	}

	<-idleConnClosed
}

//
func signalHandler(server *http.Server, idleConnClosed chan struct{},
	logger *zap.Logger) {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	<-sigChan

	logger.Info("Graceful shutdown initialised")

	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown error", zap.Error(err))
	}

	shutdownHandler(logger)

	close(idleConnClosed)
}

//
func shutdownHandler(logger *zap.Logger) {
	logger.Info("Shutdown Handler called")
}
