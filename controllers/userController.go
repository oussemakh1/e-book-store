package controllers

import (
	"log"
	"net/http"
	"rest-api/models"
	"rest-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)


type UserController struct{}

func (h UserController) CreateUser(c *gin.Context) {

	// Validate input
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return 
	}

	// Encrypt password
	err :=  input.HashPassword(input.Password)
	if err != nil {
		log.Println(err.Error())

		c.JSON(500, gin.H{
			"msg": "error hashing password",
		})
		c.Abort()

		return
	}

	// Create user
	user, err := services.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"data":user})

}


func (h UserController) FindUsers(c *gin.Context) {
	
	// Find users
	users, err := services.FindUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}


func (h UserController) FindUser(c *gin.Context) {
	
	// Get user id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return 
	}

	// Find user
	user, err := services.FindUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data":user})

}


func (h UserController) UpdateUser(c *gin.Context) {
	
	// Get user id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	// Validate input
	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	// Update user by id
	user, err := services.UpdateUser(input, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data":user})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"data": true})

}


func (h UserController) DeleteUser(c *gin.Context) {
	
	//Get user id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	// Delete user
	if err := services.DeleteUser(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"data":true})

}