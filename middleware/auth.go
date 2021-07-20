package middleware

import (
	"os"
	"rest-api/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

var theSecretKey = os.Getenv("SecretKey")

/* User with Role "READER" authentication */
func AuthReader() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")  
		if clientToken == "" {
			c.JSON(403, "No Authorization header provided")
			c.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ") 
		if len(extractedToken) == 2 {   
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(400, "Incorrect Format of Authorization Token")
			c.Abort()
			return
		}

		jwtWrapper := auth.JwtWrapper{
			SecretKey: theSecretKey,
			Issuer:    "AuthService",
		}

		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("id", claims.Id)
		c.Set("role", claims.Role)

		c.Next()

	}
}

/* User with Role "PUBLISHER" authentication */
func AuthPublisher() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(403, "No Authorization header provided")
			c.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(400, "Incorrect Format of Authorization Token")
			c.Abort()
			return
		}

		jwtWrapper := auth.JwtWrapper{
			SecretKey:       theSecretKey,
			Issuer:          "AuthService",
			ExpirationHours: 24,
		}

		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("id", claims.ID)
		c.Set("role", claims.Role)
		if claims.Role == "publisher" {
			c.Next()
		} else {
			c.JSON(401, "Not authorized")
			c.Abort()
			return
		}

	}
}

/* User with Role "ADMIN" authentication */
func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(403, "No Authorization header provided")
			c.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(400, "Incorrect Format of Authorization Token")
			c.Abort()
			return
		}

		jwtWrapper := auth.JwtWrapper{
			SecretKey:       theSecretKey,
			Issuer:          "AuthService",
		}

		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("id", claims.ID)
		c.Set("role", claims.Role)
		if claims.Role == "admin" {
			c.Next()
		} else {
			c.JSON(401, "Not authorized")
			c.Abort()
			return
		}

	}
}