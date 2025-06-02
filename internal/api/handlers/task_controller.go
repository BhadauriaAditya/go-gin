package handlers

import (
	"go-gin/internal/domain/models"
	"go-gin/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service}
}

func (t *TaskHandler) Create(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.UserID = userID
	t.service.Create(&task)
	c.JSON(http.StatusCreated, task)
}

func (t *TaskHandler) GetAll(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	tasks, _ := t.service.GetAll(userID)
	c.JSON(http.StatusOK, tasks)
}

func (t *TaskHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	userID := c.MustGet("userID").(uint)
	var task models.Task
	if c.ShouldBindJSON(&task) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input"})
		return
	}
	task.ID = uint(id)
	task.UserID = userID
	t.service.Update(&task)
	c.JSON(http.StatusOK, task)
}

func (t *TaskHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	userID := c.MustGet("userID").(uint)
	t.service.Delete(uint(id), userID)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
