package station

import (
	"context"
	"encoding/json"
	"log"

	"github.com/adopabianko/train-ticketing/database"
	"github.com/go-redis/redis/v8"
)

type IStationRepository interface {
	FindAllStationRepo() ([]Station, bool)
	FindAllStationRedisRepo() ([]Station, bool)
	CacheAllStationRepo(value interface{})
}

type StationRepository struct {
	MySQL database.IMySQLConnection
	Redis database.IRedisConnection
}

func (r *StationRepository) FindAllStationRepo() (stations []Station, status bool) {
	db := r.MySQL.CreateConnection()
	defer db.Close()

	rows, err := db.Query(`
					SELECT
						id,
					   	station_code,
					   	station_name,
					   	station_city
					FROM station
					ORDER BY station_city ASC
				`)

	defer rows.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		var st Station
		err = rows.Scan(
			&st.ID,
			&st.StationCode,
			&st.StationName,
			&st.StationCity,
		)

		if err != nil {
			log.Fatal(err.Error())
		}

		stations = append(stations, st)
	}

	if len(stations) == 0 {
		return stations, false
	}

	return stations, true
}

func (r *StationRepository) FindAllStationRedisRepo() (stations []Station, status bool) {
	cache := r.Redis.CreateConnection()

	var ctx = context.Background()
	resultCache, err := cache.Get(ctx, "station:all").Result()

	if err == redis.Nil {
		return stations, false
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal([]byte(resultCache), &stations)

	return stations, true
}

func (r *StationRepository) CacheAllStationRepo(value interface{}) {
	cache := r.Redis.CreateConnection()

	valueJson, err := json.Marshal(value)
	if err != nil {
		log.Fatal(err.Error())
	}

	var ctx = context.Background()
	err = cache.Set(ctx, "station:all", valueJson, 0).Err()

	if err != nil {
		log.Fatal(err.Error())
	}
}
