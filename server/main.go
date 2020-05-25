package server

import (
	"context"
	"time"

	"github.com/kataras/iris"
	"github.com/urfave/cli"
)

// Server used to handling a http server
type Server struct {
	IrisApp *iris.Application
	Address string
	Port    string
}

// Start server
func (s Server) Start(appContext *cli.Context) error {
	err := s.IrisApp.Run(
		iris.Addr(s.Address+":"+s.Port),
		iris.WithoutInterruptHandler)
	return err
}

// Stop server
func (s Server) Stop() error {
	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	s.IrisApp.Shutdown(ctx)
	return nil
}
