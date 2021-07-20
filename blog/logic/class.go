package logic

import (
	"blog/dao/mysql"
	"blog/models"
)
//获取分类列表
func GetClassList()([]*models.Class,error){
	return mysql.GetClassList()
}

//创建class
func CreateClass(class *models.Class)(err error){
	return mysql.CreateClass(class)
}

