package station

import (
	"github.com/adopabianko/train-ticketing/database"
)

type IStationService interface {
	StationService() (int, string, interface{})
}

type StationService struct {
	Repository IStationRepository
}

func InitStationService() *StationService {
	stationRepository := new(StationRepository)
	stationRepository.MySQL = &database.MySQLConnection{}
	stationRepository.Redis = &database.RedisConnection{}

	stationService := new(StationService)
	stationService.Repository = stationRepository

	return stationService
}

func (s *StationService) StationService() (httpCode int, message string, result interface{}) {
	resultRedis, statusRedis := s.Repository.FindAllStationRedisRepo()

	if statusRedis {
		return 200, "List of station", resultRedis
	}

	resultDB, statusDB := s.Repository.FindAllStationRepo()

	if !statusDB {
		return 404, "Empty list of station", nil
	}

	// Cache redis
	s.Repository.CacheAllStationRepo(resultDB)

	return 200, "List of station", resultDB
}
