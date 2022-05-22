package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//d, err := gorm.Open("mysql", "ahmed:gomysql@123/bookstore?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", "<username>:<password>@tcp(<host>:<port>)/<DB>?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
