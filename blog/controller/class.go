package controller

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//查询所有class
func ClassHandler(c *gin.Context){
	classes,err:=logic.GetClassList()
	if err!=nil{
		zap.L().Error("logic.GetClassList failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,classes)
}

//创建class
func CreateClass(c *gin.Context){
	//获取参数及校验
	p:=new(models.Class)
	if err:=c.ShouldBindJSON(p);err!=nil{
		zap.L().Error("CreateClass with invalid param",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	//逻辑业务处理
	if err:=logic.CreateClass(p);err!=nil{
		zap.L().Error("logic.CreateClass failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,nil)
}

