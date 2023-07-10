package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"main/dao"
	"main/models"
	"main/routers"
)

func main() {
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter()
	r.Run(":9090")
}
