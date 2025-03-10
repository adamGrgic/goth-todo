package services

import (
	"goth-todo/server/models"

	"gorm.io/gorm"
)

type TaskService interface {
	GetTasks() []models.Task
	AddTask(task *models.Task)
	ToggleTask(id string)
	DeleteTask(id string)
}

type GormTaskService struct {
	DB *gorm.DB
}

func NewGormTaskService(db *gorm.DB) *GormTaskService {
	return &GormTaskService{DB: db}
}

func (s *GormTaskService) GetTasks() []models.Task {

	var tasks []models.Task
	s.DB.Find(&tasks)
	return tasks
}

func (s *GormTaskService) AddTask(task *models.Task) {
	task.Status = "New"
	s.DB.Create(task)
}

func (s *GormTaskService) ToggleTask(id string) {
	var task models.Task
	if err := s.DB.First(&task, id).Error; err == nil {
		task.Status = "New"
		s.DB.Save(&task)
	}
}

func (s *GormTaskService) DeleteTask(id string) {
	s.DB.Delete(&models.Task{}, id)
}
