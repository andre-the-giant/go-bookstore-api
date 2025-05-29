package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"

	"go-bookstore-api/routes"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	dsn := "root:secret@tcp(127.0.0.1:3306)/bookstore"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Failed to open DB:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("❌ Failed to connect:", err)
	}

	log.Println("✅ Connected to MySQL")
}

func main() {
	initDB()

	router := gin.Default()
	routes.SetupRoutes(router, db)

	router.Run(":8080")
}
