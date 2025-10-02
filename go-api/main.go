package main

import (
    "database/sql" // Import the database/sql package
    "log"
    "net/http"
    "sync" // Import the sync package

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/go-playground/validator/v10"
    "go.uber.org/ratelimit"
)

// Declare all global variables only ONCE in this file.
var (
    db       *sql.DB
    validate = validator.New()
    cache    = sync.Map{}
)

func main() {
    // Initialize database
    InitDB("./learning_platform.db")

    // Set up Gin router
    r := gin.Default()

    // Add Sessions and Rate Limiting Middleware
    store := cookie.NewStore([]byte("secret"))
    r.Use(sessions.Sessions("mysession", store))
    rl := ratelimit.New(10) // 10 requests per second
    r.Use(func(c *gin.Context) {
        rl.Take()
        c.Next()
    })

    // Setup routes and handlers
    setupRoutes(r)

    // Serve static files
    r.StaticFS("/public", http.Dir("../templates"))

    log.Println("Go API server running on :8080")
    r.Run(":8080")
}