package server

import (
	"rest-api/controllers"
	"rest-api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Router init and enable plugins
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// CORS Configuration
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Tenant-Token", "X-Tenant-ID", "X-Tenant-User"}
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	// Gin plugins
	router.Use(gin.Recovery())

	// Controllers
	bookController := new(controllers.BookController)
	// APIs accessible publicly
	public := router.Group("/api")
	{
		public.GET("/books", bookController.FindBooks)
		public.POST("/book", bookController.CreateBook)
		public.GET("/book/:id", bookController.FindBook)
		public.PATCH("/book/:id", bookController.UpdateBook)
		public.DELETE("/book/:id", bookController.DeleteBook)
	}

	return router
}
