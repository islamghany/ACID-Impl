package api

import (
	db "github/islmaghany/bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(s db.Store) *Server {
	server := &Server{
		store: s,
	}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.getAccounts)
	router.GET("/accounts/:id", server.getAccount)
	server.router = router
	return server
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
