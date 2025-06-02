package repository

import (
	"go-crm/internal/domain/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) Create(task *models.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) GetAll(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.DB.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) Update(task *models.Task) error {
	return r.DB.Save(task).Error
}

func (r *TaskRepository) Delete(id uint, userID uint) error {
	return r.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Task{}).Error
}
