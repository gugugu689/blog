package logic

import (
	"blog/dao/redis"
	"blog/models"
	"strconv"
)
//投票
func VotePost(userID int64,p *models.ParamVote)error{
	return redis.VotePost(strconv.Itoa(int(userID)),p.PostID,float64(p.Choice))
}
