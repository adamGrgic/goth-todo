package services

import "gorm.io/gorm"

type ContentService interface {
}

type GormContentService struct {
	DB *gorm.DB
}

func NewGormContentService(db *gorm.DB) *GormContentService {
	return &GormContentService{DB: db}
}

func (s *GormContentService) GetHomePage() {

	// var tasks []models.Task
	// s.DB.Find(&tasks)
	// return tasks
}
