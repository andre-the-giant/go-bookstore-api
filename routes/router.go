package routes

import (
	"database/sql"

	"go-bookstore-api/handlers"
	"go-bookstore-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.POST("/register", handlers.Register(db))
	router.POST("/login", handlers.Login(db))
	router.POST("/logout", handlers.Logout)
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/books", handlers.GetBooks(db))
		authorized.GET("/books/:id", handlers.GetBookByID(db))
		authorized.POST("/books", handlers.PostBook(db))
		authorized.PATCH("/books/:id", handlers.UpdateQuantity(db))
		authorized.DELETE("/books/:id", handlers.DeleteBook(db))
	}
}
