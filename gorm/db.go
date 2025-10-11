package gorm

import (
	"context"
	"fmt"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
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

func SQLiteInMemoryDB(ctx context.Context) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db.Debug().WithContext(ctx)
}

func PostgreSQLDB(ctx context.Context) *gorm.DB {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=localhost user=shaouai dbname=defaultpostgresqldb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db.Debug().WithContext(ctx)
}
