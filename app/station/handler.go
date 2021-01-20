package station

import "github.com/labstack/echo"

type StationController struct {
	Service IStationService
}

func InitStationController() *StationController {
	stationService := InitStationService()

	stationController := new(StationController)
	stationController.Service = stationService
	return stationController
}

func (r *StationController) StationHandler(c echo.Context) error {
	httpCode, message, result := r.Service.StationService()

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
			"stations": result,
		},
	})
}
