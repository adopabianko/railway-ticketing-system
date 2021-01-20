package booking

import (
	"github.com/labstack/echo"
	"net/http"
)

type BookingController struct {
	Service IBookingService
}

func InitBookingController() *BookingController {
	bookingService := InitBookingService()

	bookingController := new(BookingController)
	bookingController.Service = bookingService
	return bookingController
}

func (r *BookingController)BookingHandler(c echo.Context) error{
	u := new(BookingRequest)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"code":    422,
			"message": err.Error(),
		})
	}

	httpCode, message, bookingCode := r.Service.BookingService(u.ScheduleId, u.DepartureDate, u.Qty, u.CustomerCode)

	if httpCode != 200 {
		return c.JSON(httpCode, echo.Map{
			"code":    httpCode,
			"message": message,
		})
	}

	return c.JSON(httpCode, echo.Map{
		"code": httpCode,
		"message": message,
		"data": echo.Map{
			"booking_code": bookingCode,
		},
	})
}