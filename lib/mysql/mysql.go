package mysql

import (
	"article/config"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
)

var connection *gorm.DB

func MysqlInitialize() (*gorm.DB, error) {

	dbHost := config.MysqlConfig["host"]
	dbPort := config.MysqlConfig["port"]
	dbUser := config.MysqlConfig["username"]
	dbPass := config.MysqlConfig["password"]
	dbName := config.MysqlConfig["database"]

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	var err error

	connection, err = gorm.Open("mysql", connectionString)
	if nil != err {

		redOutput := color.New(color.FgRed)
		errorOutput := redOutput.Add(color.Bold)

		errorOutput.Println("")
		errorOutput.Println(err.Error())
		errorOutput.Println("!!! Warning")
		errorOutput.Println(fmt.Sprintf("Failed connected to database %s", connectionString))
		errorOutput.Println("")

		os.Exit(2)

	} else {

		greenOutput := color.New(color.FgGreen)
		successOutput := greenOutput.Add(color.Bold)

		successOutput.Println("")
		successOutput.Println("!!! Info")
		successOutput.Println(fmt.Sprintf("Successfully connected to database %s", connectionString))
		successOutput.Println("")

	}

	fmt.Println("Connection is created")
	return connection, nil

}

func GetMysqlConnection() *gorm.DB {
	if connection == nil {
		fmt.Println("Initialize database")
		connection, _ = MysqlInitialize()
	} else {
		fmt.Println("Get connection database")
	}
	return connection
}
