package services

import "gorm.io/gorm"

type SystemService interface {
	Ping() string
}

type GormSystemService struct {
	DB *gorm.DB
}

func NewGormSystemService(db *gorm.DB) *GormTaskService {
	return &GormTaskService{DB: db}
}

func (s *GormTaskService) Ping() string {
	return "PONG"
}
