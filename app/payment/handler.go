package payment

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/xeipuuv/gojsonschema"
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

func (r *PaymentController) PaymentHandler(c echo.Context) error {
	u := new(Booking)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"code":    422,
			"message": err.Error(),
		})
	}

	valueJson, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err.Error())
	}

	schemaLoader := gojsonschema.NewReferenceLoader("file://./validation/payment_schema.json")
	documentLoader := gojsonschema.NewStringLoader(fmt.Sprintf("%s", valueJson))

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

	httpCode, message := r.Service.PaymentService(u.BookingCode)

	return c.JSON(httpCode, echo.Map{
		"code":    httpCode,
		"message": message,
	})
}
