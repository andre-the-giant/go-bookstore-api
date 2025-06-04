package main

import (
	"database/sql"
	"log"
	"os"

	"go-bookstore-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	dsn := os.Getenv("DB_DSN") // use env var
	if dsn == "" {
		log.Fatal("❌ Missing required environment variable: DB_DSN")
	}

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
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found")
	}

	initDB()

	router := gin.Default()
	routes.SetupRoutes(router, db)

	router.Run(":8080")
}
