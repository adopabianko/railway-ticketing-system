package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"log"
)

func CreateAccessToken(customerCode string) (tokenString string) {
	secretToken := viper.Get("jwt-key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access_uuid":   uuid.New(),
		"customer_code": customerCode,
	})

	tokenString, err := token.SignedString([]byte(secretToken.(string)))

	if err != nil {
		log.Fatal(err.Error())
	}

	return
}
