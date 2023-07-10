package controller

import (
	"github.com/gin-gonic/gin"
	"main/models"
	"net/http"
)

// url --> controller   -->logic -->model
// web项目的大致结构  请求来了  到控制器 到业务逻辑 到模型层的增删改查
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	//前端页面填写待办事项 点击提交 会发请求到这里

	//1.从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)

	//2.存入数据库
	err := models.CreateATodo(&todo)

	//3.返回响应
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateAtodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&todo)

	err = models.UpdateATodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteAtodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	err := models.DeleteATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			id: "deleted",
		})
	}
}
