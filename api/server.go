package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/techschool/simple_bank/db/sqlc"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// router.POST("/accounts", server.CreateAccount())

	server.router = router
	return server
}
