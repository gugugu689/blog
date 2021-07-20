package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const UserIDKey="userID"

var ErrorUserNotLogin=errors.New("请登录")

//获取当前登陆用户id
func GetCtxUserID(c *gin.Context)(userID int64,err error){
	id,ok:=c.Get(UserIDKey)
	if !ok{
		err=ErrorUserNotLogin
		return
	}
	userID=id.(int64)
	return
}
//分页
func getPageInfo(c *gin.Context) (int64, int64) {
	pageNumStr := c.Query("page")
	pageSizeStr := c.Query("size")

	var (
		pageNum int64
		pageSize int64
		err  error
	)

	pageNum, err = strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		pageNum = 1
	}
	pageSize, err = strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageSize = 10
	}
	return pageNum, pageSize
}