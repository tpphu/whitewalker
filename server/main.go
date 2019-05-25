package server

import "github.com/gin-gonic/gin"

type Server struct {
	Address string
	Port    string
}

func (self Server) Init() {

}

func (self Server) Start() {
	r := gin.Default()
	InitRouter(r)
	r.Run(self.Address + ":" + self.Port)
}
