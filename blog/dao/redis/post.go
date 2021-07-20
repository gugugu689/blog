package redis

import (
	"blog/models"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)
//发帖子
func CreatePost(postID,classID int64)(err error){
	pipeline:=client.TxPipeline()
	//发布时间
	pipeline.ZAdd(PostTimeZSet,redis.Z{
		Score: float64(time.Now().Unix()),
		Member: postID,
	})
	//分数
	pipeline.ZAdd(PostScoreZSet,redis.Z{
		Score: float64(time.Now().Unix()),
		Member: postID,
	})
	_,err= pipeline.Exec()
	return
}

//查询post_ids
func GetKeyIDs(key string,page int64,size int64)([]string,error){
	start:=(page-1)*size
	end:=start+size-1
	return client.ZRevRange(key,start,end).Result()
}

//通过匹配的order key 查询key里的ids
func GetPostIDsOrderBy(p *models.ParamPostList)([]string,error){
	keyOrder:=PostTimeZSet
	if p.Order==PostScoreZSet{
		keyOrder=PostScoreZSet
	}
	return GetKeyIDs(keyOrder,p.Page,p.Size)
}
//获取投票
func GetPostVoteYES(ids []string) (data []int64, err error) {
	// 使用pipeline一次发送多条命令,减少RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key :=PostChoiceZSetPF+id
		pipeline.ZCount(key, "1", "1")//取赞成票
	}
	cmders, err := pipeline.Exec()//cmders执行返回的统计数据
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
func GetPostVoteNO(ids []string) (data []int64, err error) {
	// 使用pipeline一次发送多条命令,减少RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key :=PostChoiceZSetPF+id
		pipeline.ZCount(key, "-1", "-1")//取赞成票
	}
	cmders, err := pipeline.Exec()//cmders执行返回的统计数据
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
//通过class获取postIDs
func GetClassPostIDsByOrder(p *models.ParamPostList)(ids []string,err error){
	KeyOrder:=PostTimeZSet
	if p.Order==PostScoreZSet{
		KeyOrder=PostScoreZSet
	}
	//zinterstore:俩zset求交集 classzset和orderzset求交集 新的交集zset：newkey
	classKey:=PostClassSetPF+strconv.Itoa(int(p.ClassID))
	newKey:=strconv.Itoa(int(p.ClassID))+KeyOrder
	//不存在newkey
	if client.Exists(newKey).Val()<1{
		pipeline:=client.Pipeline()
		pipeline.ZInterStore(newKey,redis.ZStore{
			Aggregate: "MAX",
		},classKey,KeyOrder)
		_,err=pipeline.Exec()
		if err!=nil {
			return
		}
	}
	//存在newkey
	return GetKeyIDs(newKey,p.Page,p.Size)
}
