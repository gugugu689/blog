package controller

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func VoteHandler(c *gin.Context){
	//1 参数校验
	p:=new(models.ParamVote)
	if err:=c.ShouldBindJSON(p);err!=nil{
		zap.L().Error("Vote Param false",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	//2 逻辑处理
	//     获取userID
	userID,err:=GetCtxUserID(c)
	if err!=nil{
		ResponseError(c,CodeNeedLogin)
		return
	}
	if err:=logic.VotePost(userID,p);err!=nil{
		zap.L().Error("logic.Vote failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,nil)
}
