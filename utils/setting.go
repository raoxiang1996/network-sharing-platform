package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	Dbhost     string
	DbPort     string
	Dbuser     string
	DbPassWord string
	DbName     string
)

func init() {
	cfg, err := ini.Load("../config/config.ini")
	if err != nil {
		fmt.Printf("配置文件读取错误，请检查文件路径", err)
	}
	LoadServer(cfg)
	LoadDataBase(cfg)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
}

func LoadDataBase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mongo")
	DbName = file.Section("database").Key("DbName").MustString("USTC")
	Dbhost = file.Section("database").Key("Dbhost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("27017")
	Dbuser = file.Section("database").Key("Dbuser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("raoxiang")

}
