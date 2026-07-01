package api

import (
	db "Bankcore/db/sqlc"

	"github.com/gin-gonic/gin"
)

// server serves all http request for our banking service
type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

// newServer creates a new http server and setup routing
func newServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
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
