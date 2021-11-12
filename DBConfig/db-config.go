package dbconfig

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const Dsn = "root:password@tcp(localhost)/demo?parseTime=true"

var Db *gorm.DB
var Err error

func Config() {
	Db, Err = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if Err != nil {
		panic("failed to connect database")
	}

}
