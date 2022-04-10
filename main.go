package main

import (
	"admin/app/entity"
	"admin/config/database"
	"admin/config/routes"
)

func main() {
	//连接数据库
	err := database.InitMySql()
	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库连接
	defer database.Close()
	//绑定模型
	database.SqlSession.AutoMigrate(&entity.SystemAccount{})
	//注册路由
	router := routes.SetRouter()
	//启动端口为8088的项目
	router.Run(":8088")
}
