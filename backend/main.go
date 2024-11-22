package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MICHAELKITH/Go_full_stack/config"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
    //initialize
    if err := godotenv.Load(); err != nil{
        log.Printf("Error loading from .env file: %v", err)
        return
    }

    //get our credentials from .env file
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    //build our connection 
    // Build the connection string
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
    

    //database connection 
    config.InitializeDB(dsn)

    defer config.CloseDB()
   
    app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, Backend Server")
    })

    log.Fatal(app.Listen(":3000"))
}

