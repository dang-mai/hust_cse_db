package setting

import (
	"gopkg.in/ini.v1"
)

var Conf = new(AppConfig)

// MySQLConfig 数据库配置
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

// AppConfig 运行配置
type AppConfig struct {
	Release      bool `ini:"release"`
	Port         int  `ini:"port"`
	*MySQLConfig `ini:"mysql"`
}

// Init 加载配置
func Init(file string) {
	err := ini.MapTo(Conf, file)
	if err != nil {
		panic("加载配置失败: " + err.Error())
	}
}
