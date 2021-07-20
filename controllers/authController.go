package controllers

import (
	"net/http"
	"rest-api/database"
	"rest-api/models"
	"rest-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)




func  FindMyBooks (c *gin.Context) {
	var books []models.Book
	db, _ := database.GetDB("DEV")
	author_id, _ := c.Get("id") // from the authorization middleware

	
	result := db.Table("books").Where("author_id = ?", author_id).Find(&books) // Fetch User books
	

	if result.Error == gorm.ErrRecordNotFound { // Check if user has books
		c.JSON(404, gin.H{
			"msg": "books not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": "could not get your books",
		})
		c.Abort()
		return
	}

	

	c.JSON(200, books)

	return

}


func FindMyBook(c *gin.Context) {
	var book models.Book
	db, _ := database.GetDB("DEV")
	author_id, _ := c.Get("id") // from the authorization middleware
	id, err := strconv.Atoi(c.Param("id")) // Get user id
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Table("books").Where("author_id = ? AND id = ?", author_id, id).First(&book) // Fetch the user book
	

	

	if result.Error == gorm.ErrRecordNotFound  { // Check if the user has the book
		c.JSON(404, gin.H{
			"msg": "book not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": "could not get your book",
		})
		c.Abort()
		return
	}

	

	c.JSON(200, book)

	return
}


func UpdateMyBook(c *gin.Context) {
	var verifyBook models.Book
	var input models.UpdateBookInput

	db, _ := database.GetDB("DEV")
	author_id, _ := c.Get("id") // from the authorization middleware
	id, err := strconv.Atoi(c.Param("id")) // Get the user id
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Table("books").Where("author_id = ?", author_id).First(&verifyBook) // Fetch the user book

	if result.Error == gorm.ErrRecordNotFound { //Check if the user has the book
		c.JSON(404, gin.H{
			"msg": "book not found",
		})
		c.Abort()
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil { // Validate
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := services.UpdateBook(input, id) // Update the book
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})


}

func DeleteMyBook (c *gin.Context) {
	var verify models.Book
	db, _ := database.GetDB("DEV")
	author_id, _ := c.Get("id") // from the authorization middleware
	id, err := strconv.Atoi(c.Param("id")) // Get the user id

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Table("books").Where("author_id = ?", author_id).Find(&verify) // Fetch the user book

	if result.Error == gorm.ErrRecordNotFound { // Check if the user has the book
		c.JSON(404, gin.H{
			"msg": "book not found",
		})
		c.Abort()
		return
	}
	if  err := services.DeleteBook(id); err != nil { // Delete the book
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}