package service

import (
	"github.com/Murolando/hakaton_bot_api/ent"
	"github.com/Murolando/hakaton_bot_api/pkg/repository"
)

type Actions interface {
	StopBase(tableName string) (bool, error)
	RunBase(tableName string) (bool, error)
	KillBase(sshConfig *ent.SSH) (bool, error)
	RestartBase(sshConfig *ent.SSH) error
	KillProcess(pid int64) (bool, error)
}
type Metrics interface {
	BaseStat(bdName string)(*ent.AllMetrics,error)
}
type Service struct {
	Metrics
	Actions
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Metrics: NewMetricService(repo),
		Actions: NewActionService(repo),
	}
}
