package models
//请求的参数体

type ParamSignUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"pass_word" binding:"required"`
	RePassword string `json:"re_pass_word" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVote struct {
	PostID string `json:"post_id" binding:"required"`
	Choice int8 `json:"choice" binding:"oneof=1 0 -1"`
}

const (
	Time ="time"
	Score ="score"
)
type ParamPostList struct {
	ClassID int64 `json:"class_id" form:"class_id"`
	Page int64 `json:"page" form:"page" example:"1"`
	Size int64 `json:"size" form:"size" example:"6"`
	Order string `json:"order" form:"order" example:"time"`
}