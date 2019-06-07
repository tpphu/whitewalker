package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kataras/iris"
	"github.com/urfave/cli"
)

// Server used to handling a http server
type Server struct {
	Engine  *iris.Application
	Address string
	Port    string
}

var shutdownServer = func() error { return nil }

// Start server
func (s Server) Start(appContext *cli.Context) error {
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			// kill -SIGINT XXXX or Ctrl+c
			os.Interrupt,
			syscall.SIGINT, // register that too, it should be ok
			// os.Kill  is equivalent with the syscall.Kill
			os.Kill,
			syscall.SIGKILL, // register that too, it should be ok
			// kill -SIGTERM XXXX
			syscall.SIGTERM,
		)
		select {
		case <-ch:
			timeout := 5 * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			s.Engine.Shutdown(ctx)
		}
	}()
	s.Engine.Run(iris.Addr(s.Address+":"+s.Port), iris.WithoutInterruptHandler)
	return nil
}

// Stop server
func (s Server) Stop() error {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	s.Engine.Shutdown(ctx)
	return nil
}
