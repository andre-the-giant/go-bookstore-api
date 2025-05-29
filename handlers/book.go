package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"go-bookstore-api/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title, author, quantity FROM books")
		if err != nil {
			log.Println("❌ DB Query Error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
			return
		}
		defer rows.Close()

		var books []models.Book

		for rows.Next() {
			var b models.Book
			if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity); err != nil {
				log.Println("❌ Row Scan Error:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row"})
				return
			}
			books = append(books, b)
		}

		c.IndentedJSON(http.StatusOK, books)
	}
}

func GetBookByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var b models.Book
		err = db.QueryRow("SELECT id, title, author, quantity FROM books WHERE id = ?", id).
			Scan(&b.ID, &b.Title, &b.Author, &b.Quantity)

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query book"})
			return
		}

		c.IndentedJSON(http.StatusOK, b)
	}
}

func PostBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newBook models.Book

		if err := c.BindJSON(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		if newBook.Title == "" || newBook.Author == "" || newBook.Quantity < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid fields"})
			return
		}

		// Insert into DB
		result, err := db.Exec("INSERT INTO books (title, author, quantity) VALUES (?, ?, ?)",
			newBook.Title, newBook.Author, newBook.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert book"})
			return
		}

		insertedID, err := result.LastInsertId()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inserted ID"})
			return
		}
		newBook.ID = int(insertedID)

		c.IndentedJSON(http.StatusCreated, newBook)
	}
}

func DeleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		result, err := db.Exec("DELETE FROM books WHERE id = ?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	}
}

func UpdateQuantity(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var input struct {
			Quantity int `json:"quantity"`
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if input.Quantity < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be 0 or greater"})
			return
		}

		result, err := db.Exec("UPDATE books SET quantity = ? WHERE id = ?", input.Quantity, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Quantity updated"})
	}
}
