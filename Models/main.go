package models

import (
	configdb "github.com/hzprog/restapi/DBConfig"
	Book "github.com/hzprog/restapi/Models/book"
)

func InitilizeMigation() {
	//migrate the schemes
	configdb.Db.AutoMigrate(Book.Book{})
}
