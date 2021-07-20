package server

import (
	"rest-api/auth"
	"rest-api/controllers"
	"rest-api/docs"
	"rest-api/middleware"

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
	userController := new(controllers.UserController)
	api := router.Group("/api")
	{
					// APIs accessible publicly
			public := api.Group("/")
			{	
				// login api
				public.POST("/login",auth.Login)

			}
			// APIs accessible by reader
			reader := api.Group("/reader").Use(middleware.AuthReader())
			{
				// Fetch all books 
				reader.GET("/books", bookController.FindBooks)
				// Fetch book by id 
				reader.GET("/book/:id", bookController.FindBook)

			}

			// APIs accessible by publisher
			publisher := api.Group("/publisher").Use(middleware.AuthPublisher())
			{
				// Fetch all books 
				publisher.GET("/books", bookController.FindBooks)
				// Fetch book by id 
				publisher.GET("/book/:id", bookController.FindBook)
				// Fetch current publisher created books
				publisher.GET("/mybooks",controllers. FindMyBooks)
				// Fetch current publisher book by id
				publisher.GET("/mybook/:id", controllers.FindMyBook)
				// Create  book
				publisher.POST("/createbook", bookController.CreateBook)
				// Updated current publisher book by id
				publisher.PATCH("/mybook/:id", controllers.UpdateMyBook)
				// Delete current publisher book by id
				publisher.DELETE("/mybook/:id", controllers.DeleteMyBook)

			}

			// APIs accessible by admin 
			admin := api.Group("/admin").Use(middleware.AuthAdmin())
			{

			/* Books REST-API*/
				// Fetch all books 
				admin.GET("/books", bookController.FindBooks)
				// Create  book
				admin.POST("/book", bookController.CreateBook)
				// Fetch book by id 
				admin.GET("/book/:id", bookController.FindBook)
				// Update book by id 
				admin.PATCH("/book/:id", bookController.UpdateBook)
				// Delete book by id 
				admin.DELETE("/book/:id", bookController.DeleteBook)

			
			/* Users REST-API*/
				// Fetch all users
				admin.GET("/users", userController.FindUsers)
				// Create user
				admin.POST("/user", userController.CreateUser)
				// Fetch user by id
				admin.GET("/user/:id", userController.FindUser)
				// Update user by id
				admin.PATCH("/user/:id", userController.UpdateUser)
				// Delete user by id
				admin.DELETE("/user/:id", userController.DeleteUser)	
			}

	}

	

	return router
}
