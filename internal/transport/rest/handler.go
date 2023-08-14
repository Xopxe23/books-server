package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/xopxe23/books-server/docs"
	"github.com/xopxe23/books-server/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swag/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/refresh", h.refresh)
	}
	books := router.Group("/books", h.userIdentity)
	{
		books.POST("/", h.createBook)
		books.GET("/", h.getAllBooks)
		books.GET("/:id", h.getBookById)
		books.PUT("/:id", h.updateBook)
		books.DELETE("/:id", h.deleteBook)
	}
	return router
}
