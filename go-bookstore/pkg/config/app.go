package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "sql10722458:HnXjbwVfsd@tcp(sql10.freemysqlhosting.net:3306)/sql10722458?charset=utf8&parseTime=True&loc=UTC")
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
