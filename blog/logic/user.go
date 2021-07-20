package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"blog/pkg/jwt"
	"blog/pkg/snowflake"

)
//注册
func SignUp(p *models.ParamSignUp)(err error){
	//1 判断用户是否存在
	if err:=mysql.CheckUserExist(p.Username);err!=nil{
		return err
	}
	//2 雪花生成userID
	userID:=snowflake.GenID()
	//3 构造User实例 将ParamSignUp数据放入User
	user:=&models.User{
		UserID: userID,
		Username: p.Username,
		Password: p.Password,
	}
	//4 放入数据库
	return mysql.InsertUser(user)
}
//登陆
func Login(p *models.ParamLogin)(user *models.User,err error){
	user=&models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err=mysql.Login(user);err!=nil{
		return
	}
	token,err1:=jwt.GenToken(user.UserID,user.Username)
	if err1!=nil{
		return
	}
	user.Token=token
	return
}
