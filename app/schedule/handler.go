package schedule

import (
	"github.com/labstack/echo"
)

type ScheduleController struct {
	Service IScheduleService
}

func InitScheduleController() *ScheduleController {
	scheduleService := InitScheduleService()

	scheduleController := new(ScheduleController)
	scheduleController.Service = scheduleService
	return scheduleController
}

func (r *ScheduleController) ScheduleHandler(c echo.Context) error {
	origin := c.QueryParam("origin")
	destination := c.QueryParam("destination")
	departureDate := c.QueryParam("departure_date")

	httpCode, message, result := r.Service.ScheduleService(origin, destination, departureDate)

	if httpCode != 200 {
		return c.JSON(httpCode, echo.Map{
			"code":    httpCode,
			"message": message,
		})
	}

	return c.JSON(httpCode, echo.Map{
		"code":    httpCode,
		"message": message,
		"data": echo.Map{
			"schedules": result,
		},
	})
}
