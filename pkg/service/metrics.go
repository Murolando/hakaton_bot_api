package service

import "github.com/Murolando/hakaton_bot_api/pkg/repository"

type MetricService struct {
	repo *repository.Repository
}

func NewMetricService(repo *repository.Repository) *MetricService {
	return &MetricService{
		repo: repo,
	}
}
