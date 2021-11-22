package models

import (
	configdb "github.com/hzprog/restapi/DBConfig"
	Book "github.com/hzprog/restapi/Models/book"
	User "github.com/hzprog/restapi/Models/user"
)

func InitializeMigration() {
	//migrate the schemes
	configdb.Db.AutoMigrate(Book.Book{}, User.User{})
}
