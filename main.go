package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

// 声明一个数据库的名字
var (
	DB *gorm.DB
)

// 用来在数据库中描述一个待办事项
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 连接数据库
func initMySQL() (err error) {
	//连接数据库
	dsn := "root:20020902=QWer@(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	//注意这里的err不要去声明 用我们的函数参数
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	//这里我们测试一下能否连接到 如果不能会返回相应错误
	return DB.DB().Ping()
}
func main() {
	//连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	//模型绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	//告诉gin框架如果需要加载静态文件 去哪里找
	r.Static("/static", "static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//建立一个组 用来处理同类业务
	v1Group := r.Group("v1")
	{
		//添加事项
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面填写待办事项 点击提交 会发请求到这里
			//1.从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			//2.存入数据库
			err = DB.Create(&todo).Error
			//3.返回响应
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//查看所有待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			err = DB.Find(&todoList).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		//查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改某一个待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "无效的id",
				})
				return
			}
			var todo Todo
			err = DB.Where("id = ?", id).First(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			}
			c.BindJSON(&todo)
			err = DB.Save(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//删除某一个待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "无效的id",
				})
				return
			}
			err = DB.Where("id=?", id).Delete(Todo{}).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					id: "deleted",
				})
			}
		})
	}
	r.Run(":9090")
}
