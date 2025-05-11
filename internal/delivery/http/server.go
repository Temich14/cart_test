package http

import (
	"context"
	_ "github.com/Temich14/cart_test/docs"
	"github.com/Temich14/cart_test/internal/config"
	"github.com/Temich14/cart_test/internal/delivery/http/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"
)

type Server struct {
	Server *http.Server
	api    *gin.Engine
	cfg    *config.ServerConfig
	logger *slog.Logger
}

func NewServer(cfg *config.ServerConfig, logger *slog.Logger) *Server {
	return &Server{
		api:    gin.New(),
		cfg:    cfg,
		logger: logger,
	}
}
func (s *Server) RegisterHandlers(registerFunc func(engine *gin.RouterGroup), groupURL string) {
	registerFunc(s.api.Group(groupURL))
}
func (s *Server) registerMiddleware(middleware ...gin.HandlerFunc) {
	s.api.Use(middleware...)
}
func createCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	})
}
func (s *Server) Run() error {
	s.logger.Info("starting server")
	s.initDocs()
	s.registerMiddleware(middleware.TokenClaimer(s.cfg.Secret), createCors())
	s.Server = &http.Server{
		Addr:    ":" + s.cfg.Port,
		Handler: s.api,
	}
	err := s.Server.ListenAndServe()
	if err != nil {
		s.logger.Error("error starting server", slog.String("err", err.Error()), slog.String("stack", string(debug.Stack())))
		return err
	}
	s.logger.Info("server started and listen on " + s.Server.Addr)
	return nil
}
func (s *Server) Stop(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
func (s *Server) initDocs() {
	s.api.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})
	s.api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
