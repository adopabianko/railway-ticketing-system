package main

import (
	"github.com/adopabianko/train-ticketing/routes"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".app-config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	routes.Routes()
}
