package main

import (
	"fmt"
	"hust/cse/db/models"
	"hust/cse/db/mysql"
	"hust/cse/db/router"
	"hust/cse/db/setting"
)

func main() {
	// 加载配置文件
	setting.Init("config/config.ini")
	// 连接数据库
	mysql.InitMySQL(setting.Conf.MySQLConfig)
	// 绑定模型
	migrate()
	// 注册路由
	r := router.SetupRouter()
	// 运行
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("Server startup failed, err: %v\n", err)
	}
	// 关闭数据库
	defer mysql.CloseMySQL()
}

// Migrate 迁移
func migrate() {
	var err error
	err = mysql.DB.AutoMigrate(&models.Student{})
	if err != nil {
		panic("迁移失败: " + err.Error())
	}
	err = mysql.DB.AutoMigrate(&models.Course{})
	if err != nil {
		panic("迁移失败: " + err.Error())
	}
	err = mysql.DB.AutoMigrate(&models.SC{})
	if err != nil {
		panic("迁移失败: " + err.Error())
	}
}
