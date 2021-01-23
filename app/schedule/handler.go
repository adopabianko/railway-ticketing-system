package schedule

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/xeipuuv/gojsonschema"
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

	schemaLoader := gojsonschema.NewReferenceLoader("file://./validation/schedule_schema.json")

	m := map[string]interface{}{
		"origin":         origin,
		"destination":    destination,
		"departure_date": departureDate,
	}
	documentLoader := gojsonschema.NewGoLoader(m)

	validate, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		log.Fatal(err.Error())
	}

	if !validate.Valid() {
		for _, desc := range validate.Errors() {
			return c.JSON(419, echo.Map{
				"code":    419,
				"message": fmt.Sprintf("%s", desc),
			})
		}
	}

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
