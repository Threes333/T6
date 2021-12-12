package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db  *gorm.DB //数据库句柄
	err error    //错误信息
)

// @title 初始化数据库
// @description 初始化与数据库的连接与设置
func InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/threes?parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	//将模型与数据库连接
	if err := db.AutoMigrate(User{}, Message{}); err != nil {
		panic(err)
	}
}
