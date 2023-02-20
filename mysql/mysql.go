package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"hust/cse/db/setting"
)

var (
	DB  *gorm.DB
	err error
)

// InitMySQL 连接数据库并初始化配置
func InitMySQL(cfg *setting.MySQLConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	DB, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gorm_",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}
}

// CloseMySQL 关闭数据库
func CloseMySQL() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic("关闭数据库连接失败: " + err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		panic("关闭数据库连接失败: " + err.Error())
	}
}
