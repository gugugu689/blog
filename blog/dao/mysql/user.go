package mysql

import (
	"blog/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

const secret="the first work"

//检查用户是否存在
func CheckUserExist(username string)(err error){
	sqlstr:=`select count(user_id) from user where username=?`
	var count int64
	if err:=db.Get(&count,sqlstr,username);err!=nil{
		return err
	}
	if count>0{
		return ErrorUserExist
	}
	return
}
//密码加密
func encryptPassword(password string)string{
	h:=md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
//插入用户
func InsertUser(user *models.User)(err error){
	//密码加密
	user.Password=encryptPassword(user.Password)
	sqlstr:=`insert into user(user_id,username,password) values(?,?,?)`
	_,err=db.Exec(sqlstr,user.UserID,user.Username,user.Password)
	return
}
//登陆
func Login(user *models.User)(err error){
	sqlstr:=`select username,password from user where username=?`
	inputPassword:=user.Password
	err=db.Get(user,sqlstr,user.Username)
	if err==sql.ErrNoRows{
		return ErrorUserNotExist
	}
	if err!=nil{
		return err
	}
	sqlPassword:=encryptPassword(inputPassword)
	if sqlPassword!=user.Password{
		return  ErrorInvalidPassword
	}
	return
}
//根据id取用户
func GetUser(userID int64)(user *models.User,err error){
	user=new(models.User)
	sqlstr:=`select user_id,user_name from user where user_id=?`
	err=db.Get(user,sqlstr,userID)
	return
}
