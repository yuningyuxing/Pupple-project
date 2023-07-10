package models

import (
	"main/dao"
)

// 用来在数据库中描述一个待办事项
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//Todo 这个model 的增删改查操作都放在这里

func CreateATodo(todo *Todo) (err error) {
	//2.存入数据库
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todoList *[]Todo, err error) {
	todoList = new([]Todo)
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	err = dao.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
