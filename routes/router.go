package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"go-bookstore-api/handlers"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	router.Use(loggerMiddleware())

	router.GET("/books", handlers.GetBooks(db))
	router.GET("/books/:id", handlers.GetBookByID(db))
	router.POST("/books", handlers.PostBook(db))
	router.PATCH("/books/:id", handlers.UpdateQuantity(db))
	router.DELETE("/books/:id", handlers.DeleteBook(db))
}

func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path

		// Log method and path
		println("➡️", method, path)

		c.Next() // Pass on to the next handler

		status := c.Writer.Status()
		println("⬅️", status)
	}
}
