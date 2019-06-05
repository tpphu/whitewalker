package server

import (
	"context"
	"log"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

// Server used to handling a http server
type Server struct {
	Engine  *gin.Engine
	Address string
	Port    string
}

var shutdownServer = func() error { return nil }

// Start server
func (s Server) Start(appContext *cli.Context) error {
	endPoint := s.Address + ":" + s.Port
	readTimeout := time.Second * 60
	writeTimeout := time.Second * 60
	maxHeaderBytes := 1 << 10
	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes
	server := endless.NewServer(
		endPoint,
		s.Engine)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	shutdownServer = func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return server.Shutdown(ctx)
	}

	return server.ListenAndServe()
}

// Stop server
func (s Server) Stop() error {
	return shutdownServer()
}
