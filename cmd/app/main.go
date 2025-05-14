package main

import (
	"context"
	"github.com/Temich14/cart_test/internal/app"
	"github.com/Temich14/cart_test/internal/config"
	"github.com/Temich14/cart_test/internal/logger"
	"github.com/Temich14/cart_test/internal/migrator"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title						Cart API
// @version					1.0
// @description				API для управления корзиной пользователя и его заказами
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @host						localhost:8080
// @BasePath
func main() {
	cfg := config.MustLoad()
	migrator.NewMigrator("migrations/", cfg.DBConfig.Conn).MustApplyMigrations()
	log := logger.New(cfg.Env)
	log.Info("starting server")
	application := app.NewApp(cfg, log)

	application.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop // graceful shutdown

	log.Info("shutting down application")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	application.Stop(ctx)

	log.Info("application successfully stopped")
}
