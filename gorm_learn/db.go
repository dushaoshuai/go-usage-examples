package gorm_learn

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DefaultGormDB() *gorm.DB {
	dsn := "ingtube:Quqwpe_940301@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
