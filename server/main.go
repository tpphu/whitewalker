package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Server used to handling a http server
type Server struct {
	Engine  *gin.Engine
	Address string
	Port    string
}

var srv *http.Server

// Start server
func (s Server) Start() error {
	addr := s.Address + ":" + s.Port
	srv = &http.Server{
		Addr:    addr,
		Handler: s.Engine,
	}
	return srv.ListenAndServe()
}

// Stop server
func (s Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}
