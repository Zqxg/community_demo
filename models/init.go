package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

// 获得一个*grom.DB对象
var DB *gorm.DB

// 连接数据库
func Database(pathstring string) {
	// 链接Mysql数据库
	db, err := gorm.Open(mysql.Open(pathstring), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})

	if err != nil {
		log.Println(err)
		log.Panic("mysql数据库连接错误")
	}

	fmt.Println("数据库连接成功")

	// 设置数据库连接池
	sqlDb, err := db.DB()
	if err != nil {
		log.Println("获取数据库连接池失败：", err)
		log.Panic("获取数据库连接池失败")
	}

	// 输出连接池状态
	fmt.Println("连接池状态：", sqlDb.Stats())

	// 设置连接池参数
	sqlDb.SetMaxIdleConns(10)                  // 设置连接池中空闲连接的最大数量
	sqlDb.SetMaxOpenConns(20)                  // 设置数据库的最大连接数量
	sqlDb.SetConnMaxLifetime(30 * time.Second) // 设置连接可复用的最大时间

	DB = db
	// 数据库迁移
	DB.AutoMigrate(&Topic{}, &Post{})
}
