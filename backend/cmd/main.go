package main

import (
    "os"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/joho/godotenv"
    "log"
    "backend/internal/handlers"
)

func main() {
    // Load env vars
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORS())

    // OAuth2 Google login
    e.GET("/auth/google/login", handlers.GoogleLogin)
    e.GET("/auth/google/callback", handlers.GoogleCallback)

    // JWT-protected API
    api := e.Group("/api")
    api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
        SigningKey: []byte(os.Getenv("JWT_SECRET")),
    }))
    api.POST("/teams", handlers.CreateTeam)
    api.POST("/projects", handlers.CreateProject)
    api.GET("/spec/:id", handlers.GetSpec)
    api.PUT("/spec/:id", handlers.UpdateSpec)

    // WebSocket (with JWT auth)
    e.GET("/ws", handlers.WSHandler)

    // Health check
    e.GET("/healthz", func(c echo.Context) error { return c.String(200, "ok") })

    e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
} 