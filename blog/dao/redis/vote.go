package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"time"
)
//投票
func VotePost(userID string,postID string,choice float64)error{
	//1 判断投票限制
	//    redis取发布时间
	postTime:=client.ZScore(PostTimeZSet,postID).Val()
	if float64(time.Now().Unix())-postTime>7*24*3600{
		return errors.New("投票时间已过")
	}
	//2 更新帖子分数
	//    查看当前用户投票记录
	oc:=client.ZScore(PostChoiceZSetPF+postID,userID).Val()
	if oc==choice{
		return errors.New("重复投票")
	}
	diff:=choice-oc
	pipeline:=client.TxPipeline()
	pipeline.ZIncrBy(PostScoreZSet,diff,postID)
	//3 记录(添加)用户投票数据
	if choice==0{
		pipeline.ZRem(PostChoiceZSetPF+postID,userID)
	}else{
		pipeline.ZAdd(PostChoiceZSetPF+postID,redis.Z{
			Score: choice,
			Member: userID,
		})
	}
	_,err:=pipeline.Exec()
	return err
}