package rest

import (
    "github.com/gin-gonic/gin"
    "backend/go-gin/api/rest/handler"
    "backend/go-gin/internal/middleware"
)

func InitRoutes() *gin.Engine {
    r := gin.Default()

    // Global middleware
    r.Use(middleware.RequestLogger())
    r.Use(middleware.AuthOptional()) // Public middleware, fallback for now

    // Health route (no auth)
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // Group: /api/v1 (Protected)
    api := r.Group("/api/v1")
    api.Use(middleware.AuthRequired()) // Protected routes

    user := api.Group("/users")
    {
        user.POST("/", handler.CreateUser)
        user.GET("/", handler.ListUsers)
        user.GET("/:id", handler.GetUser)
        user.PUT("/:id", handler.UpdateUser)
        user.DELETE("/:id", handler.DeleteUser)
    }

    return r
}
