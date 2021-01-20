package schedule

import (
	"github.com/adopabianko/train-ticketing/database"
)

type IScheduleService interface {
	ScheduleService(org string, des string, depDat string) (int, string, interface{})
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
	resultRedis, statusRedis := s.Repository.FindAllScheduleRedisRepo()

	if statusRedis {
		return 200, "List of schedule", resultRedis
	}

	resultDB, statusDB := s.Repository.FindAllScheduleRepo(org, des, depDate)

	if !statusDB {
		return 404, "Empty list of schedule", nil
	}

	// Cache redis
	s.Repository.CacheAllScheduleRepo(resultDB)

	return 200, "List of schedule", resultDB
}
