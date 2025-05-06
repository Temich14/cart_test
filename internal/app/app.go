package app

import (
	"context"
	"github.com/Temich14/cart_test/internal/config"
	"github.com/Temich14/cart_test/internal/delivery/http"
	"github.com/Temich14/cart_test/internal/delivery/http/handler/cart"
	order2 "github.com/Temich14/cart_test/internal/delivery/http/handler/order"
	"github.com/Temich14/cart_test/internal/repository"
	cart2 "github.com/Temich14/cart_test/internal/service/cart"
	"github.com/Temich14/cart_test/internal/service/order"
	"log"
	"log/slog"
	"runtime/debug"
)

type App struct {
	server *http.Server
	repo   DBCloser
	cfg    *config.AppConfig
	logger *slog.Logger
}

func NewApp(cfg *config.AppConfig, logger *slog.Logger) *App {
	return &App{cfg: cfg, logger: logger}
}
func (a *App) Run() {
	a.server = http.NewServer(a.cfg.ServerConfig, a.logger)

	repo := repository.NewRepository(a.cfg.DBConfig, a.logger, a.cfg.Env)

	a.repo = repo

	cartService := cart2.NewCartService(repo, a.logger)
	cartHandler := cart.NewHandler(cartService, a.logger)

	orderService := order.NewOrderService(repo, a.logger)
	orderHandler := order2.NewHandler(orderService, a.logger)

	a.server.RegisterHandlers(cartHandler.Register, "cart/")
	a.server.RegisterHandlers(orderHandler.Register, "order/")

	go func() {
		err := a.server.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
func (a *App) Stop(ctx context.Context) {
	err := a.server.Stop(ctx)
	if err != nil {
		a.logger.Error("error stopping server", slog.String("error", err.Error()), slog.String("stack", string(debug.Stack())))
	}
	err = a.repo.CloseDB()
	if err != nil {
		a.logger.Error("error closing db", slog.String("error", err.Error()), slog.String("stack", string(debug.Stack())))
	}
}
