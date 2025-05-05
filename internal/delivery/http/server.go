package http

import (
	"context"
	_ "github.com/Temich14/cart_test/docs"
	"github.com/Temich14/cart_test/internal/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

type Server struct {
	Server *http.Server
	api    *gin.Engine
	cfg    *config.ServerConfig
}

func NewServer(cfg *config.ServerConfig) *Server {
	return &Server{
		api: gin.New(),
		cfg: cfg,
	}
}
func (s *Server) RegisterHandlers(registerFunc func(engine *gin.RouterGroup), groupURL string) {
	registerFunc(s.api.Group(groupURL))
}
func (s *Server) Run() error {
	s.initDocs()
	s.Server = &http.Server{
		Addr:    ":" + s.cfg.Port,
		Handler: s.api,
	}
	return s.Server.ListenAndServe()
}
func (s *Server) Stop(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
func (s *Server) initDocs() {
	s.api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
