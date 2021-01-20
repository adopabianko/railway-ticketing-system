package payment

import (
	"github.com/labstack/echo"
	"net/http"
)

type PaymentController struct {
	Service IPaymentService
}

func InitPaymentController() *PaymentController {
	paymentService := InitPaymentService()

	paymentController := new(PaymentController)
	paymentController.Service = paymentService
	return paymentController
}

func(r *PaymentController)PaymentHandler(c echo.Context)error{
	u := new(Booking)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"code":    422,
			"message": err.Error(),
		})
	}

	httpCode, message := r.Service.PaymentService(u.BookingCode)

	return c.JSON(httpCode, echo.Map{
		"code": httpCode,
		"message": message,
	})
}
