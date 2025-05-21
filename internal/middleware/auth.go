package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

const validToken = "mysecrettoken" // You can load from ENV later

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")

        if token != "Bearer "+validToken {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "unauthorized",
            })
            return
        }

        // continue
        c.Next()
    }
}

// Optional auth (for public routes or future enhancement)
func AuthOptional() gin.HandlerFunc {
    return func(c *gin.Context) {
        // No auth enforced here, but can inspect headers
        c.Next()
    }
}
