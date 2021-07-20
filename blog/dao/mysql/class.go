package mysql

import (
	"blog/models"
	"database/sql"
	"go.uber.org/zap"
)
//获取分类列表
func GetClassList()(ClassList []*models.Class,err error){
	ClassList=make([]*models.Class,0,5)
	sqlstr:=`select class_id,class_name from class`
	err=db.Select(&ClassList,sqlstr)
	if err==sql.ErrNoRows{
		zap.L().Warn("no class in")
		err=nil
	}
	if err!=nil{
		zap.L().Error("select class failed",zap.Error(err))
		return
	}
	return
}
//创建class
func CreateClass(class *models.Class)(err error){
	sqlstr:=`insert into class(class_name) values(?)`
	_,err=db.Exec(sqlstr,class.Name)
	return
}
//根据class_id获取class
func GetClass(classID int64)(class *models.Class,err error){
	class=new(models.Class)
	sqlstr:=`select class_ld,class_name from class where class_id=?`
	err=db.Get(class,sqlstr,classID)
	return
}
