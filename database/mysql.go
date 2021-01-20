package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type IMySQLConnection interface {
	CreateConnection() *sql.DB
}

type MySQLConnection struct{}

func (c *MySQLConnection) CreateConnection() (db *sql.DB) {
	db, err := sql.Open("mysql", fmt.Sprintf("%v", viper.Get("mysql")))
	if err != nil {
		log.Fatal(err.Error())
	}

	return
}
