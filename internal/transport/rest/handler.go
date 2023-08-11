package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/xopxe23/books-server/docs"
	"github.com/xopxe23/books-server/internal/domain"
)

type Books interface {
	Create(book domain.Book) (int, error)
	GetAll() ([]domain.Book, error)
	GetById(id int) (domain.Book, error)
	Update(id int, input domain.UpdateBookInput) error
	Delete(id int) error
}

type Handler struct {
	bookService Books
}

func NewHandler(books Books) *Handler {
	return &Handler{bookService: books}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swag/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

// @Summary Create Book
// @Tags Books REST
// @ID create-book
// @Accept json
// @Produce json
// @Param input body domain.Book true "book info"
// @Success 200 {integer} integer
// @Failure default {object} errorResponse
// @Router /books/ [post]
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

// @Summary Get All Books
// @Tags Books REST
// @ID get-all-books
// @Accept json
// @Produce json
// @Success 200 {integer} []domain.Book
// @Failure default {object} errorResponse
// @Router /books/ [get]
func (h Handler) getAllBooks(c *gin.Context) {
	books, err := h.bookService.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": books,
	})
}

// @Summary Get Book by ID
// @Tags Books REST
// @ID get-book-by-id
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {integer} domain.Book
// @Failure default {object} errorResponse
// @Router /books/{id} [get]
func (h Handler) getBookById(c *gin.Context) {
	var book domain.Book
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	book, err = h.bookService.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)
}

// @Summary Update Book
// @Tags Books REST
// @ID update-book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param input body domain.UpdateBookInput true "Update input"
// @Success 200 {integer} string
// @Failure default {object} errorResponse
// @Router /books/{id} [put]
func (h Handler) updateBook(c *gin.Context) {
	var input domain.UpdateBookInput
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = h.bookService.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

// @Summary Delete Book
// @Tags Books REST
// @ID delete-book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {integer} string
// @Failure default {object} errorResponse
// @Router /books/{id} [delete]
func (h Handler) deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = h.bookService.Delete(id); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
