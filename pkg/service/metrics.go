package service

import (
	"github.com/Murolando/hakaton_bot_api/ent"
	"github.com/Murolando/hakaton_bot_api/pkg/repository"
)

type MetricService struct {
	repo *repository.Repository
}

func NewMetricService(repo *repository.Repository) *MetricService {
	return &MetricService{
		repo: repo,
	}
}

func (s *MetricService) BaseStat(bdName string) (*ent.AllMetrics, error) {
	dataBaseStat,err := s.repo.Metrics.DataBaseStat()
	if err!=nil{
		return nil,err
	}
	pgActivityStat,err := s.repo.Metrics.PgActivityStat()
	if err!=nil{
		return nil,err
	}
	pgDatabaseSize,err := s.repo.Metrics.PgDatabaseSize(bdName)
	if err!=nil{
		return nil,err
	}
	metrics := ent.AllMetrics{
		DataBaseStat: *dataBaseStat,
		PgActivityStat: *pgActivityStat,
		PgDatabaseSize: *pgDatabaseSize,
	}
	return &metrics,nil
}
