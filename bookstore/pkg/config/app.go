package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() error {
	DB, err := gorm.Open("mysql", "saikumar:Saikumar@3093@/simplerest?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		return err
	}
	db = DB
	return nil

}

func GetDB() *gorm.DB {
	return db
}
