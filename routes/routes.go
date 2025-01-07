package routes

import (
	"book-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	route := gin.Default()

	route.POST("/books", controllers.CreateBook)
	route.GET("/books", controllers.GetBooks)
	route.GET("/books/:id", controllers.GetBookByID)
	route.PUT("/books/:id", controllers.UpdateBook)
	route.DELETE("/books/:id", controllers.DeleteBook)

	return route
}
