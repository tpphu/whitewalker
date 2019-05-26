package server

import (
	"github.com/gin-gonic/gin"
)

// Server used to handling a http server
type Server struct {
	Engine  *gin.Engine
	Address string
	Port    string
}

func (self Server) Init() {

}

func (self Server) Start() {
	self.Engine.Run(self.Address + ":" + self.Port)
}
