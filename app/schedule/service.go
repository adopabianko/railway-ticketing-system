package schedule

import (
	"github.com/adopabianko/train-ticketing/database"
)

type IScheduleService interface {
	ScheduleService(org, des, depDate string) (int, string, interface{})
}

type ScheduleService struct {
	Repository IScheduleRepository
}

func InitScheduleService() *ScheduleService {
	scheduleRepository := new(ScheduleRepository)
	scheduleRepository.MySQL = &database.MySQLConnection{}
	scheduleRepository.Redis = &database.RedisConnection{}

	scheduleService := new(ScheduleService)
	scheduleService.Repository = scheduleRepository

	return scheduleService
}

func (s *ScheduleService) ScheduleService(org string, des string, depDate string) (httpCode int, message string, result interface{}) {
	resultRedis, statusRedis := s.Repository.FindScheduleRedisRepo(org, des, depDate)

	if statusRedis {
		return 200, "List of schedule", resultRedis
	}

	resultDB, statusDB := s.Repository.FindScheduleRepo(org, des, depDate)

	if !statusDB {
		return 404, "Empty list of schedule", nil
	}

	// Cache redis
	s.Repository.CacheScheduleRepo(org, des, depDate, resultDB)

	return 200, "List of schedule", resultDB
}
