package controller

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)
//写评论
func AddCommentHandler(c *gin.Context){
	//1 获取参数绑定和校验
	p:=new(models.Comment)
	if err:=c.ShouldBindJSON(p);err!=nil{
		zap.L().Error("create comment whith invalid param",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	// 获取post_id
	sPostID:=c.Param("id")
	PostID,err:=strconv.ParseInt(sPostID,10,64)
	if err!=nil{
		zap.L().Error("AddComment with incalid param",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	// 获取user_id
	userID,err:=GetCtxUserID(c)
	if err!=nil{
		ResponseError(c,CodeNeedLogin)
		return
	}
	p.PostID=PostID
	p.UserID=userID
	//2 业务逻辑处理
	if err:=logic.AddComment(p);err!=nil{
		zap.L().Error("logic.AddComment failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	//3 返回响应
	ResponseSuccess(c,nil)
}
//获取评论
func GetComment(c *gin.Context){
	//取post_id
	sid:=c.Param("id")
	post_id,err:=strconv.ParseInt(sid,10,64)
	if err!=nil{
		zap.L().Error("GetPost id failed",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	comments,err:=logic.GetCommentByPostID(post_id)
	if err!=nil{
	zap.L().Error("logic.GetCommentByPostID failed",zap.Error(err))
	ResponseError(c,CodeServerBusy)
	return
	}
	ResponseSuccess(c,comments)
}
