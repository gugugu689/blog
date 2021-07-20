package models

/*comment_id bigint  auto_increment primary key,
post_id bigint not null,
content varchar(8192) not null,
user_id bigint not null,
*/

type Comment struct {
	ID int64 `json:"id" db:"comment_id"`
	PostID int64 `json:"post_id" db:"post_id"`
	Content string `json:"content" db:"content" binding:"required"`
	UserID int64 `json:"user_id" db:"user_id"`
}