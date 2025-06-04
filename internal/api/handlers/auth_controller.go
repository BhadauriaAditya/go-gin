package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"go-gin/internal/domain/models"
	"go-gin/internal/infrastructure/database"
	"go-gin/internal/domain/repository"	
	"go-gin/internal/service"
)

type AuthRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := &models.User{Email: req.Email, Password: string(hash)}

	db, _ := database.GetDB("gin")
	if err := repository.CreateUser(db, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db, _ := database.GetDB("gin")
	user, err := repository.GetUserByEmail(db, req.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := service.GenerateJWT(user.Email)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
