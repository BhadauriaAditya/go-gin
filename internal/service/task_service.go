package service

import (
	"go-gin/internal/domain/models"
	"go-gin/internal/domain/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo}
}

func (s *TaskService) Create(task *models.Task) error {
	return s.repo.Create(task)
}

func (s *TaskService) GetAll(userID uint) ([]models.Task, error) {
	return s.repo.GetAll(userID)
}

func (s *TaskService) Update(task *models.Task) error {
	return s.repo.Update(task)
}

func (s *TaskService) Delete(id, userID uint) error {
	return s.repo.Delete(id, userID)
}
