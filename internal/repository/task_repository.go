package repository

import (
	"goth-todo/internal/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.DB.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) AddTask(task *models.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) ToggleTask(id string) error {
	var task models.Task
	err := r.DB.First(&task, "id = ?", id).Error
	if err != nil {
		return err
	}
	return r.DB.Save(&task).Error
}

// func (r *TaskRepository) DeleteTask(id string) {
// 	r.DB.Delete(&models.Task{}, id)
// }
