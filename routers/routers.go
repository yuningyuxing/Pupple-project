package routers

import (
	"github.com/gin-gonic/gin"
	"main/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	//告诉gin框架如果需要加载静态文件 去哪里找
	r.Static("/static", "static")
	r.GET("/", controller.IndexHandler)

	//建立一个组 用来处理同类业务
	v1Group := r.Group("v1")
	{
		//添加事项
		v1Group.POST("/todo", controller.CreateTodo)

		//查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)

		//修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateAtodo)

		//删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteAtodo)
	}
	return r
}
