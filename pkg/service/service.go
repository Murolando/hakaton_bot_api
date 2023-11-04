package service

import (
	"github.com/Murolando/hakaton_bot_api/pkg/repository"
)

type Metrics interface{

}
type Service struct {
	Metrics
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Metrics: NewMetricService(repo),
	}
}
