package model

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"University-Information-Website/utils"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.Dbuser,
		utils.DbPassWord,
		utils.Dbhost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		fmt.Printf("连接数据库失败,请检查参数: ", err)
	}
	// 禁用表的复数形式
	db.SingularTable(true)

	//db.AutoMigrate()

	// 设置连接池最大闲置连接数
	db.DB().SetMaxIdleConns(10)
	// 设置数据库最大连接数
	db.DB().SetMaxOpenConns(100)
	// 设置最大可复用时间
	db.DB().SetConnMaxLifetime(10 * time.Second)
	//db.Close()
}
