package controller

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)
//写帖子
func CreatePostHandler(c *gin.Context){
	//1 获取参数及校验
	p:=new(models.Post)
	if err:=c.ShouldBindJSON(p);err!=nil{
		zap.L().Error("create post failed",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	// 获取用户id
	userID,err:=GetCtxUserID(c)
	if err!=nil{
		ResponseError(c,CodeNeedLogin)
		return
	}
	p.AuthorID=userID
	//2 业务处理
	if err:=logic.CreatePost(p);err!=nil{
		zap.L().Error("logic.CreatePost failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,nil)
}
//看帖子
func GetPostHandler(c *gin.Context){
	//取post_id
	sid:=c.Param("id")
	post_id,err:=strconv.ParseInt(sid,10,64)
	if err!=nil{
		zap.L().Error("GetPost id failed",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	PostDetail,err:=logic.GetPostDetail(post_id)
	if err!=nil{
		zap.L().Error("logic.GetPost failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,PostDetail)
}
//获取帖子列表
func GetPostListHandler(c *gin.Context){
	pageNum, pageSize:= getPageInfo(c)
	posts,err:=logic.GetPostList(pageNum,pageSize)
	if err!=nil{
		zap.L().Error("logic.GetPostList failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,posts)
}
// 根据 分数 时间 获取帖子列表
func GetPostListBy_Handler(c *gin.Context){
	//1 获取参数 校验
	//      GET (query string)(url里的)
	//      初始化models.ParamPostList
	p:=&models.ParamPostList{
		Page: 1,
		Size: 6,
		Order: models.Time,
	}
	if err:=c.ShouldBindQuery(p);err!=nil{
		zap.L().Error(" ParamPostList shouldbindquery failed",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	//2 逻辑处理
	postlist,err:=logic.GetPostListBy_(p)
	if err!=nil{
		zap.L().Error("logic.GetPostListBy failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,postlist)
}
//按照class获取帖子列表
func GetPostListByClassIDHandler(c *gin.Context){
	//1 获取参数 校验
	//      GET (query string)(url里的)
	//      初始化models.ParamPostList
	p:=&models.ParamPostList{
		Page: 1,
		Size: 6,
		Order: models.Time,
	}
	if err:=c.ShouldBindQuery(p);err!=nil{
		zap.L().Error(" ParamPostList shouldbindquery failed",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	//2 逻辑处理
	postlist,err:=logic.GetPostListByClassID(p)
	if err!=nil{
		zap.L().Error("logic.GetPostListByClassID failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,postlist)
}