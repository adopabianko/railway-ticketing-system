package booking

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/xeipuuv/gojsonschema"
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

func (r *BookingController) BookingHandler(c echo.Context) error {
	u := new(BookingRequest)
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

	schemaLoader := gojsonschema.NewReferenceLoader("file://./validation/booking_schema.json")
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

	// Save booking data
	httpCode, message, bookingCode := r.Service.BookingService(u)

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
			"booking_code": bookingCode,
		},
	})
}

func (r *BookingController) FindBookingDetail(c echo.Context) error {
	csCode := c.QueryParam("customer_code")
	bookingCode := c.QueryParam("booking_code")

	schemaLoader := gojsonschema.NewReferenceLoader("file://./validation/booking_detail_schema.json")

	m := map[string]interface{}{
		"customer_code": csCode,
		"booking_code":  bookingCode,
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

	httpCode, message, result := r.Service.FindBookingDetailService(csCode, bookingCode)

	if httpCode != 200 {
		return c.JSON(httpCode, echo.Map{
			"code":    httpCode,
			"message": message,
		})
	}

	return c.JSON(httpCode, echo.Map{
		"code":    httpCode,
		"message": message,
		"data":    result,
	})
}
