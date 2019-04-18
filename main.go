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

	"github.com/josnidhin/go-rest-api-example/handlers"
	"github.com/josnidhin/go-rest-api-example/system"
)

const (
	ShutdownTimeout = 10 * time.Second
)

func main() {
	app := &system.App{}
	server := &http.Server{}
	idleConnClosed := make(chan struct{})

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)
		signal.Notify(sigChan, syscall.SIGTERM)

		<-sigChan

		app.Logger.Info("Graceful shutdown initialised")

		ctx, _ := context.WithTimeout(context.Background(), ShutdownTimeout)

		if err := server.Shutdown(ctx); err != nil {
			app.Logger.Error("Server shutdown error", zap.Error(err))
		}

		close(idleConnClosed)
	}()

	app.Config = system.ConfigInstance()
	app.Logger = system.LoggerInstance(app.Config)
	app.DB = system.NewPGDB(app.Config)

	handlers.New(app)
	app.Router = system.NewRouter(Routes(), handlers.Default404)

	serverAddress := fmt.Sprintf(":%d", app.Config.Server.HTTP.Port)

	server.Addr = serverAddress
	server.Handler = app.Router

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		app.Logger.Fatal("Server startup failed", zap.Error(err))
	}

	<-idleConnClosed
}
