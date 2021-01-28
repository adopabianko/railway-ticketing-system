package routes

import (
	"net/http"

	"github.com/adopabianko/train-ticketing/app/auth"
	"github.com/adopabianko/train-ticketing/app/booking"
	"github.com/adopabianko/train-ticketing/app/payment"
	"github.com/adopabianko/train-ticketing/app/schedule"
	"github.com/adopabianko/train-ticketing/app/station"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

type JwtCustomClaims struct {
	AccessUid    string `json:"access_uid"`
	CustomerCode string `json:"customer_code"`
	jwt.StandardClaims
}

func Routes() {
	authController := auth.InitAuthController()
	stationController := station.InitStationController()
	scheduleController := schedule.InitScheduleController()
	bookingController := booking.InitBookingController()
	paymentController := payment.InitPaymentController()

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	secretKey := viper.Get("jwt-key")
	configJwt := middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(secretKey.(string)),
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"app":     "Railway Ticketing System",
			"version": "0.0.1",
		})
	})

	e.POST("/auth/register", authController.RegisterHandler)
	e.POST("/auth/activation", authController.ActivationHandler)
	e.POST("/auth/login", authController.LoginHandler)
	e.GET("/station", stationController.StationHandler, middleware.JWTWithConfig(configJwt))
	e.GET("/schedule", scheduleController.ScheduleHandler, middleware.JWTWithConfig(configJwt))
	e.POST("/booking", bookingController.BookingHandler, middleware.JWTWithConfig(configJwt))
	e.GET("/booking/detail", bookingController.FindBookingDetail, middleware.JWTWithConfig(configJwt))
	e.POST("/payment", paymentController.PaymentHandler, middleware.JWTWithConfig(configJwt))

	e.Logger.Fatal(e.Start(":3000"))
}
