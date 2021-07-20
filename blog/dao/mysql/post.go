package mysql

import (
	"blog/models"
	"github.com/jmoiron/sqlx"
	"strings"
)
//发帖子
func CreatePost(post *models.Post)(err error){
	sqlstr:=`insert into post(post_id,title,content,author_id,class_id) values(?,?,?,?,?)`
	_,err=db.Exec(sqlstr,post.ID,post.Title,post.Content,post.AuthorID,post.ClassID)
	return
}
//看帖子
func GetPost(post_id int64)(post *models.Post,err error){
	post=new(models.Post)
	sqlstr:=`select post_id,author_id,class_id,title,content,create_time from post where post_id=?`
	err=db.Get(post,sqlstr,post_id)
	return
}
//获取帖子列表
func GetPostList(pageNum int64,pageSize int64)(posts []*models.Post,err error){
	sqlstr:=`select post_id,author_id,class_id,title,content,create_time from post order by create_time desc limit ?,?`
	posts=make([]*models.Post,0,2)
	err=db.Select(posts,sqlstr,pageNum,pageSize)
	if err!=nil{
		return
	}
	return
}
//ids查询posts
func GetPostListByIDs(ids []string)(posts []*models.Post,err error){
	//FIND_IN_SET按照参数传进去的顺序进行查询
	sqlstr:=`select post_id,author_id,class_id,title,content,create_time from post where post_id in (?) order by FIND_IN_SET(post_id,?)`
	//sqlx.in 方便用于sql in函数
	//query 返回的新question语句   args 返回的新参数
	query,args,err:=sqlx.In(sqlstr,ids,strings.Join(ids,","))
	if err!=nil {
		return
	}
	//将question语句 query 翻译为sql语句 query
	query=db.Rebind(query)
	err=db.Select(&posts,query,args...)//args参数不为单一 需...
	return
}