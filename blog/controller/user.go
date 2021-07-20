package controller

import (
	"blog/dao/mysql"
	"blog/logic"
	"blog/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

)
//注册
func SignUpHandler(c *gin.Context){
	//1 参数获取和校验 （从前端拿到param shouldbindjson后 放入ParamSignUp）
	p:=new(models.ParamSignUp)
	if err:=c.ShouldBindJSON(p);err!=nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2 业务逻辑处理 （处理ParamSignUp）
	if err:=logic.SignUp(p);err!=nil{
		zap.L().Error("logic.SingnUp failde",zap.Error(err))
		if errors.Is(err,mysql.ErrorUserExist){
			ResponseError(c,CodeUserExist)
			return
		}
		ResponseError(c,CodeServerBusy)
		return
	}
	//3 返回响应
	ResponseSuccess(c,nil)
}
//登陆
func LoginHandler(c *gin.Context){
	//获取请求参数和校验
	p:=new(models.ParamLogin)
	if err:=c.ShouldBindJSON(p);err!=nil{
		zap.L().Error("login with invalid param",zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//业务逻辑处理
	user,err:=logic.Login(p)
	if err!=nil{
		zap.L().Error("logic.login failed",zap.Error(err))
		if errors.Is(err,mysql.ErrorUserNotExist){
			ResponseError(c,CodeUserNotExist)
			return
		}
		ResponseError(c,CodeInvalidPassword)
		return
	}
	//返回响应
	ResponseSuccess(c,gin.H{
		"username": user.Username,
		"token": user.Token,
	})
}
