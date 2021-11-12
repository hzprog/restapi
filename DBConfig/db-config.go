package dbconfig

import (
	"fmt"

	Env "github.com/hzprog/restapi/Helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var username string = Env.GetEnvVar("DATABASE_USERNAME")
var password string = Env.GetEnvVar("DATABASE_PASSWORD")
var url string = Env.GetEnvVar("DATABASE_URL")
var port string = Env.GetEnvVar("DATABASE_PORT")
var dbname string = Env.GetEnvVar("DATABASE_NAME")

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, url, port, dbname)

var Db *gorm.DB
var Err error

func Config() {
	Db, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if Err != nil {
		panic("failed to connect database")
	}
}
