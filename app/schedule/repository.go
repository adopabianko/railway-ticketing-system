package schedule

import (
	"context"
	"encoding/json"
	"log"

	"github.com/adopabianko/train-ticketing/database"
	"github.com/go-redis/redis/v8"
)

type IScheduleRepository interface {
	FindAllScheduleRepo(string, string, string) ([]Schedule, bool)
	FindAllScheduleRedisRepo() ([]Schedule, bool)
	CacheAllScheduleRepo(value interface{})
}

type ScheduleRepository struct {
	MySQL database.IMySQLConnection
	Redis database.IRedisConnection
}

func (r *ScheduleRepository) FindAllScheduleRepo(org, des, depDate string) (schedules []Schedule, status bool) {
	db := r.MySQL.CreateConnection()
	defer db.Close()

	rows, err := db.Query(`
		SELECT
			id,
			origin,
			destination,
			train_code,
			time,
			quota,
			balance,
			price,
			start_date,
			end_date
		FROM schedule
		WHERE origin = ? 
			AND destination = ?
			AND status_active = 1
		ORDER BY start_date, time ASC
	`, org, des)

	defer rows.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		var sc Schedule
		err := rows.Scan(
			&sc.ID,
			&sc.Origin,
			&sc.Destination,
			&sc.TrainCode,
			&sc.Time,
			&sc.Quota,
			&sc.Balance,
			&sc.Price,
			&sc.StartDate,
			&sc.EndDate,
		)

		if err != nil {
			log.Fatal(err.Error())
		}

		if depDate >= sc.StartDate && depDate <= sc.EndDate {
			schedules = append(schedules, sc)
		}
	}

	if len(schedules) == 0 {
		return schedules, false
	}

	return schedules, true
}

func (r *ScheduleRepository) FindAllScheduleRedisRepo() (schedules []Schedule, status bool) {
	cache := r.Redis.CreateConnection()

	var ctx = context.Background()
	resultCache, err := cache.Get(ctx, "schedule:all").Result()

	if err == redis.Nil {
		return schedules, false
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal([]byte(resultCache), &schedules)

	return schedules, true
}

func (r *ScheduleRepository) CacheAllScheduleRepo(value interface{}) {
	cache := r.Redis.CreateConnection()

	valueJson, err := json.Marshal(value)
	if err != nil {
		log.Fatal(err.Error())
	}

	var ctx = context.Background()
	err = cache.Set(ctx, "schedule:all", valueJson, 0).Err()

	if err != nil {
		log.Fatal(err.Error())
	}
}
