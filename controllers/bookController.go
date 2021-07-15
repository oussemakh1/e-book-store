package controllers

import (
	"net/http"
	"rest-api/models"
	"rest-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct{}

// CreateBook godoc
// @Summary Create new book
// @Description Create new book
// @Security Authorization
// @Accept  json
// @Produce  json
// @Param  input body models.CreateBookInput true "book data"
// @Success 200 {string} string	"ok"
// @Router /api/book [Post]
func (h BookController) CreateBook(c *gin.Context) {

	// Validate input
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book, err := services.CreateBook(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// FindBooks godoc
// @Summary Get all books
// @Description Get all books
// @Security Authorization
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /api/books [Get]
func (h BookController) FindBooks(c *gin.Context) {

	// find books
	books, err := services.FindBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// FindBook godoc
// @Summary Find a book
// @Description Find a book
// @Security Authorization
// @Produce  json
// @Param id path int true "book identifier"
// @Success 200 {string} string	"ok"
// @Router /api/book/:id [Get]
func (h BookController) FindBook(c *gin.Context) {
	//Get book id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find book
	book, err := services.FindBook(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update a book
// @Security Authorization
// @Accept  json
// @Produce  json
// @Param id path int true "book identifier"
// @Param  input body models.CreateBookInput true "book data"
// @Success 200 {string} string	"ok"
// @Router /api/book/:id [Patch]
func (h BookController) UpdateBook(c *gin.Context) {

	//Get book id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update book by id
	book, err := services.UpdateBook(input, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book
// @Security Authorization
// @Produce  json
// @Param id path int true "book identifier"
// @Success 200 {string} string	"ok"
// @Router /api/book/:id [Delete]
func (h BookController) DeleteBook(c *gin.Context) {
	//Get book id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.DeleteBook(id); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
