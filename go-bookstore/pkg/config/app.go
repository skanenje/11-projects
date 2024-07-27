package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)
func Connect(){
	d, err := gorm.Open("mysql", "sql7722456:qk5ZmEa8Tc@sql7.freemysqlhosting.net:3306/sql7722456?charset=utf8&parseTime=True&loc=local")
	if err != nil {
		log.Fatal(err)
	}
	db = d
}
func GetDb()*gorm.DB{
	return db
}