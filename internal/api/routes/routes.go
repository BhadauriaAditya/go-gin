package routes

import (
	"go-crm/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(r *gin.Engine, authHandler *handlers.AuthHandler) {
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.GET("/users", authHandler.Users)
}

func ProtectedRoutes(r *gin.RouterGroup, taskHandler *handlers.TaskHandler) {
	r.POST("/tasks", taskHandler.Create)
	r.GET("/tasks", taskHandler.GetAll)
	r.PUT("/tasks/:id", taskHandler.Update)
	r.DELETE("/tasks/:id", taskHandler.Delete)
}
