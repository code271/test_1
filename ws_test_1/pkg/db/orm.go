package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

type Option func(*Config)

type Config struct {
	DataSourceName string
	MaxIdleConn    int
	MaxOpenConn    int
	MaxLifetime    time.Duration
	PrefixMapper   string
	LogMode        bool
}

func Init(c Option) (err error) {
	mc := new(Config)
	c(mc)
	DB, err = gorm.Open("mysql", mc.DataSourceName)
	if err != nil {
		return
	}
	DB.DB().SetMaxIdleConns(mc.MaxIdleConn)
	DB.DB().SetMaxOpenConns(mc.MaxOpenConn)
	DB.DB().SetConnMaxLifetime(mc.MaxLifetime)
	DB.LogMode(mc.LogMode)
	return
}

func Close() {
	if DB != nil {
		_ = DB.Close()
	}
}
