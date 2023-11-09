package api

import (
	"time"

	"github.com/gin-gonic/gin"
	schema "github.com/techschool/simplebank/api/schema"
	db "github.com/techschool/simplebank/db/sqlc"
)

// Server server HTTP requests for our banking service
type Server struct {
	store *db.Store
	route *gin.Engine
}

func NewServer(sotre *db.Store) *Server {
	server := &Server{store: sotre}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/account/:id", server.getAccountById)
	router.GET("/accounts/", server.getListAccount)

	server.route = router
	return server
}

// Start run the HTTP server a specific address
func (server *Server) Start(address string) error {
	return server.route.Run(address)
}

func errResponse(err error, requestId string, requestTime time.Time) gin.H {
	return gin.H{"error": schema.GetResponseError(err, requestId, requestTime)}
}
