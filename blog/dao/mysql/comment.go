package mysql

import (
	"blog/models"
	"go.uber.org/zap"
)
//添加评论
func AddComment(comment *models.Comment)(err error){
	sqlstr:=`insert into comment(post_id,content,user_id) values(?,?,?)`
	_,err=db.Exec(sqlstr,comment.PostID,comment.Content,comment.UserID)
	if err!=nil{
		return
	}
	return
}
//获取评论
func GetCommentByPostID(post_id int64)(comments []*models.Comment,err error){
	comments=make([]*models.Comment,0,5)
	sqlstr:=`select content,user_id from comment`
	if err=db.Select(&comments,sqlstr);err!=nil{
		zap.L().Error("select comments failed",zap.Error(err))
		return
	}
	return
}