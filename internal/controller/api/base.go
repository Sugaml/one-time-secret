package api

import (
	"fmt"

	_ "github.com/berrybytes/simplesecrets/docs"
	"github.com/berrybytes/simplesecrets/internal/controller/token"
	db "github.com/berrybytes/simplesecrets/internal/model/sqlc"
	"github.com/berrybytes/simplesecrets/util"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot make token maker : %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil

}
func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/", server.heathCheck)
	router.GET("/healthcheck", server.heathCheck)
	router.POST("/user", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/one-time-secret", server.createOneTimeSecret)
	router.GET("/one-time-secret/:id/:content", server.getOneTimeSecret)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/secrets", server.createSecret)
	authRoutes.GET("/secrets/:id/:content", server.getSecret)
	authRoutes.GET("/secrets", server.listSecrets)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
