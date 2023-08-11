package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xopxe23/books-server/internal/domain"
)

type Books interface {
	Create(book domain.Book) (int, error)
}

type Handler struct {
	bookService Books
}

func NewHandler(books Books) *Handler {
	return &Handler{bookService: books}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	books := router.Group("/books")
	{
		books.POST("/", h.createBook)
		books.GET("/", h.getAllBooks)
		books.GET("/:id", h.getBookById)
		books.PUT("/:id", h.updateBook)
		books.DELETE("/:id", h.deleteBook)
	}
	return router
}

func (h Handler) createBook(c *gin.Context) {
	var book domain.Book
	if err := c.Bind(&book); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.bookService.Create(book)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h Handler) getAllBooks(c *gin.Context) {}
func (h Handler) getBookById(c *gin.Context) {}
func (h Handler) updateBook(c *gin.Context)  {}
func (h Handler) deleteBook(c *gin.Context)  {}
