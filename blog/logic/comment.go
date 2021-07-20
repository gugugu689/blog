package logic

import (
	"blog/dao/mysql"
	"blog/models"
)
//添加评论
func AddComment(comment *models.Comment)(err error){
	if err=mysql.AddComment(comment);err!=nil{
		return
	}
	return
}
//获取评论
func GetCommentByPostID(post_id int64)(comments []*models.Comment,err error){
	return mysql.GetCommentByPostID(post_id)
}