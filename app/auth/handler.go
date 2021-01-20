package auth

import (
	"github.com/adopabianko/train-ticketing/jwt"
	"net/http"

	"github.com/labstack/echo"
)

type AuthController struct {
	Service IAuthService
}

func InitAuthController() *AuthController {
	authService := InitAuthService()

	authController := new(AuthController)
	authController.Service = authService
	return authController
}

func (r *AuthController) RegisterHandler(c echo.Context) error {
	u := new(Customer)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"code":    422,
			"message": err.Error(),
		})
	}

	httpCode, message, result := r.Service.RegisterService(u)

	if httpCode != 200 {
		return c.JSON(httpCode, echo.Map{
			"code": httpCode,
			"message": message,
		})
	}

	return c.JSON(httpCode, echo.Map{
		"code": httpCode,
		"message": message,
		"data": result,
	})
}

func (r *AuthController)ActivationHandler(c echo.Context) error {
	u := new(Customer)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"code":    422,
			"message": err.Error(),
		})
	}

	httpCode, message := r.Service.ActivationService(u.CustomerCode, u.ActivationCode)

	return c.JSON(httpCode, echo.Map{
		"code": httpCode,
		"message": message,
	})
}

func(r *AuthController)LoginHandler(c echo.Context) error {
	u := new(Customer)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"code":    422,
			"message": err.Error(),
		})
	}

	httpCode, message, result := r.Service.LoginService(u.Email, u.Password)

	if httpCode != 200 {
		return c.JSON(httpCode, echo.Map{
			"code": httpCode,
			"message": message,
		})
	}

	accessToken := jwt.CreateAccessToken(result.CustomerCode)

	return c.JSON(httpCode, echo.Map{
		"code": httpCode,
		"message": message,
		"data": echo.Map{
			"user": echo.Map{
				"id": result.ID,
				"customer_code": result.CustomerCode,
				"first_name": result.FirstName,
				"last_name": result.LastName,
				"email": result.Email,
				"phone_number": result.PhoneNumber,
				"gender": result.Gender,
				"birth_date": result.BirthDate,
				"activation_code": result.ActivationCode,
			},
			"access_token": accessToken,
		},
	})
}
