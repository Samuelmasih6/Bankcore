package api

import (
	db "github.com/Samuelmasih6/Bankcore/db/sqlc"

	"github.com/gin-gonic/gin"
)

// server serves all http request for our banking service
type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

// newServer creates a new http server and setup routing
func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	server.router = router
	return server
}

// start run the http on specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
