package dbconfig

import (
	Env "github.com/hzprog/restapi/Helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var username string = Env.GetEnvVar("DATABASE_USERNAME")
var password string = Env.GetEnvVar("DATABASE_PASSWORD")
var url string = Env.GetEnvVar("DATABASE_URL")
var port string = Env.GetEnvVar("DATABASE_PORT")
var dbname string = Env.GetEnvVar("DATABASE_NAME")

var dsn string = username + ":" + password + "@tcp(" + url + ":" + port + ")/" + dbname + "?parseTime=true"

var Db *gorm.DB
var Err error

func Config() {
	Db, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if Err != nil {
		panic("failed to connect database")
	}
}
