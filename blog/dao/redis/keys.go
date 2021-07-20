package redis

const(
	PostTimeZSet="post:time"
	PostScoreZSet="post:score"
	PostChoiceZSetPF="post:choice:"  //+post_id  key:userID value:choice
	PostClassSetPF="class:" //+post_id key:class_id value:post_id
)