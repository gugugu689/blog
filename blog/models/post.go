package models

import "time"

type Post struct {
	ID int64 `json:"id,string" db:"post_id"`
	AuthorID int64 `json:"author_id" db:"author_id"`
	ClassID int64 `json:"class_id" db:"class_id" binding:"required"`
	Title string `json:"title" db:"title" binding:"required"`
	Content string `json:"content" db:"content" binding:"required"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

type PostDetail struct {
	VoteYES int64 `json:"vote_yes"`
	VoteNO int64 `json:"vote_no"`
	*Post
	AuthorName string
	*Class	`json:"class"`
}
