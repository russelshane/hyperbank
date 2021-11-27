package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/russelshane/hyperbank/db/sqlc"
	"github.com/russelshane/hyperbank/token"
	"github.com/russelshane/hyperbank/util"
)

// Server will serve HTTP requests for hyperbank service
type Server struct {
	config util.Config
	store db.Store
	tokenHandler token.Maker
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	token, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token handler: %w", err)
	}

	server := &Server{
		config: config,
		store: store,
		tokenHandler: token,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

// Setup API routes and handlers
func (server *Server) setupRouter() {
	router := gin.Default()

	// Bank Transfers routes
	router.POST("/api/v8/transfers", server.makeTransfer)

	// User routes
	router.POST("/api/v8/users", server.createUser)
	router.POST("/api/v8/users/login", server.loginUser)

	// Account routes
	router.POST("/api/v8/accounts", server.createAccount)
	router.PUT("/api/v8/accounts/update", server.updateAccount)
	router.GET("/api/v8/accounts/:id", server.getAccount)
	router.GET("/api/v8/accounts", server.listAccounts)

	server.router = router
}

// Start runs an HTTP on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}