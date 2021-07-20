package logic

import (
	"blog/dao/mysql"
	"blog/dao/redis"
	"blog/models"
	"blog/pkg/snowflake"
	"go.uber.org/zap"
)
//发帖子
func CreatePost(post *models.Post)(err error){
	//生成post_id
	post.ID=snowflake.GenID()
	if err=mysql.CreatePost(post);err!=nil{
		return
	}
	if err=mysql.CreatePost(post);err!=nil{
		return
	}
	return
}
//查看帖子
func GetPostDetail(post_id int64)(PostDetail *models.PostDetail,err error){
	post,err:=mysql.GetPost(post_id)
	if err!=nil{
		zap.L().Error("mysql.GetPost failed",zap.Error(err))
		return
	}
	user,err:=mysql.GetUser(post.AuthorID)
	if err!=nil{
		zap.L().Error("mysql.GetUser failed",zap.Error(err))
		return
	}
	class,err:=mysql.GetClass(post.ClassID)
	if err!=nil{
		zap.L().Error("mysql.GetClass failed",zap.Error(err))
		return
	}
	PostDetail=&models.PostDetail{
		AuthorName: user.Username,
		Post: post,
		Class: class,
	}
	return
}
//获取帖子列表
func GetPostList(pageNum int64,pageSize int64)(data []*models.PostDetail,err error){
	posts,err1:=mysql.GetPostList(pageNum,pageSize)
	if err1!=nil{
		zap.L().Error("mysql.GetPostList failed",zap.Error(err1))
		return
	}
	data=make([]*models.PostDetail,0,len(posts))
	for _,post:=range posts{
		user,err2:=mysql.GetUser(post.AuthorID)
		if err2!=nil{
			zap.L().Error("mysql.GetUser failed",zap.Error(err2))
			return
		}
		class,err3:=mysql.GetClass(post.ClassID)
		if err3!=nil{
			zap.L().Error("mysql.GetClass failed",zap.Error(err3))
			return
		}
		postdetail:=&models.PostDetail{
			AuthorName: user.Username,
			Post: post,
			Class: class,
		}
		data=append(data, postdetail)
	}
	return
}
//按照（分数/时间）排名 获取帖子列表
func GetPostListBy_(p *models.ParamPostList)(postlist []*models.PostDetail,err error){
	//1 redis 查询post_ids
	ids,err:=redis.GetPostIDsOrderBy(p)
	if err!=nil{
		zap.L().Error("redis.GetPostIDsOrderBy failed",zap.Error(err))
		return
	}
	//2 mysql 根据post_ids查posts
	posts,err:=mysql.GetPostListByIDs(ids)
	if err!=nil{
		zap.L().Error("mysql.GetPostListByIDs failed",zap.Error(err))
		return
	}
	//3 查询帖子投票数
	voteYES,err:=redis.GetPostVoteYES(ids)
	if err!=nil{
		zap.L().Error("redis.GetPostVoteYES failed",zap.Error(err))
		return
	}
	voteNO,err:=redis.GetPostVoteNO(ids)
	if err!=nil{
		zap.L().Error("redis.GetPostVoteNO failed",zap.Error(err))
		return
	}
	//4 组合数据 user class post
	postlist=make([]*models.PostDetail,0,len(posts))
	for index,post:=range posts{
		user,err1:=mysql.GetUser(post.AuthorID)
		if err1!=nil{
			zap.L().Error("mysql.GetUser failed",zap.Error(err1))
			return
		}
		class,err2:=mysql.GetClass(post.ClassID)
		if err2!=nil{
			zap.L().Error("mysql.GetClass failed",zap.Error(err2))
			return
		}
		postdetail:=&models.PostDetail{
			AuthorName: user.Username,
			Post: post,
			Class: class,
			VoteYES: voteYES[index],
			VoteNO: voteNO[index],
		}
		postlist=append(postlist, postdetail)
	}
	return
}
//通过class获取帖子列表
func GetPostListByClassID(p *models.ParamPostList)(postlist []*models.PostDetail,err error){
	//1 redis查询post_ids
	ids,err:=redis.GetClassPostIDsByOrder(p)
	if err!=nil{
		zap.L().Error("redis.GetClassPostsByOrder failed",zap.Error(err))
		return
	}
	//2 根据ids mysql查posts
	posts,err:=mysql.GetPostListByIDs(ids)
	if err!=nil{
		zap.L().Error("mysql.GetPostListByIDs failed",zap.Error(err))
		return
	}
	//3 查询好每post的vote
	voteYES,err:=redis.GetPostVoteYES(ids)
	if err!=nil{
		zap.L().Error("redis.GetPostVoteYES failed",zap.Error(err))
		return
	}
	voteNO,err:=redis.GetPostVoteNO(ids)
	if err!=nil{
		zap.L().Error("redis.GetPostVoteNO failed",zap.Error(err))
		return
	}
	//4 组合数据
	postlist=make([]*models.PostDetail,0,len(posts))
	for index,post:=range posts{
		user,err1:=mysql.GetUser(post.AuthorID)
		if err1!=nil{
			zap.L().Error("mysql.GetUser failed",zap.Error(err1))
			return
		}
		class,err2:=mysql.GetClass(post.ClassID)
		if err2!=nil{
			zap.L().Error("mysql.GetClass failed",zap.Error(err2))
			return
		}
		postdetail:=&models.PostDetail{
			AuthorName: user.Username,
			Post: post,
			Class: class,
			VoteYES: voteYES[index],
			VoteNO: voteNO[index],
		}
		postlist=append(postlist, postdetail)
	}
	return
}
