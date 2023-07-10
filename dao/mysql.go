package dao

import "github.com/jinzhu/gorm"

// 声明一个数据库名字
var (
	DB *gorm.DB
)

// 连接数据库
func InitMySQL() (err error) {
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

func Close() {
	DB.Close()
}
