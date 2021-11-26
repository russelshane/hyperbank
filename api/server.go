package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/russelshane/hyperbank/db/sqlc"
)

// Server will serve HTTP requests for hyperbank service
type Server struct {
	store db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/api/v8/accounts", server.createAccount)
	router.POST("/api/v8/transfers", server.makeTransfer)
	router.PUT("/api/v8/accounts/update", server.updateAccount)
	router.GET("/api/v8/accounts/:id", server.getAccount)
	router.GET("/api/v8/accounts", server.listAccounts)

	server.router = router
	return server
}

// Start runs an HTTP on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}