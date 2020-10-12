package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"os"
)

func CreateConnection() (*gorm.DB, error) {
	//host := os.Getenv("DB_HOST")
	//user := os.Getenv("DB_USER")
	//DBName := os.Getenv("DB_NAME")
	//password := os.Getenv("DB_PASSWORD")
	host := "127.0.0.1:3308"
	user := "root"
	DBName := "micro-homestead"
	password := "secret"

	return gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, host, DBName,
		),
	)
}
