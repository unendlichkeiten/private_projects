package conf

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

const (
	_dsn = "root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
)

func init() {
	var err error
	DB, err = NewMysql()
	if err != nil {
		log.Fatalf("new mysql err [%s]", err)
		return
	}
}

func NewMysql() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: _dsn}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// set connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
