package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"iiujapp.tech/basic-gin/conf"
	"iiujapp.tech/basic-gin/handler"
	"iiujapp.tech/basic-gin/service"
)

// Server struct
type Server struct {
	srv *http.Server
	s   service.Iservice
}

// NewServer function
func NewServer(s service.Iservice) *Server {
	server := &Server{
		s: s,
	}
	return server
}

// Start function
func (server *Server) Start() {
	server.ginStart()
}

func (server Server) ginStart() {

	r := server.SetupRouter()
	r.Run(conf.ServerPort)
}

// SetupRouter function
func (server Server) SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		handler.UserHandler(server.s, c)
	})

	r.POST("/user", func(c *gin.Context) {
		handler.UserSaveHandler(server.s, c)
	})

	return r
}
