package gorm_learn

import (
	"context"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DefaultGormDB(ctx context.Context) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("mysql_user"),
		os.Getenv("mysql_passwd"),
		os.Getenv("mysql_host"),
		os.Getenv("mysql_port"),
		os.Getenv("mysql_db_name"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db.Debug().WithContext(ctx)
}
