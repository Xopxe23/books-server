package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xopxe23/books-server/internal/domain"
)

// @Summary Create Book
// @Security ApiKeyAuth
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
	id, err := h.services.Books.Create(book)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All Books
// @Security ApiKeyAuth
// @Tags Books REST
// @ID get-all-books
// @Accept json
// @Produce json
// @Success 200 {integer} []domain.Book
// @Failure default {object} errorResponse
// @Router /books/ [get]
func (h Handler) getAllBooks(c *gin.Context) {
	books, err := h.services.Books.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": books,
	})
}

// @Summary Get Book by ID
// @Security ApiKeyAuth
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
	book, err = h.services.Books.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)
}

// @Summary Update Book
// @Security ApiKeyAuth
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
	if err = h.services.Books.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

// @Summary Delete Book
// @Security ApiKeyAuth
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
	if err = h.services.Books.Delete(id); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
