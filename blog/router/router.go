package router
import (
	"blog/controller"
	"blog/logger"
	"blog/middlewares"




	"github.com/gin-gonic/gin"
)

func SetupRouter()*gin.Engine{
	r:=gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))


	v1:=r.Group("/blog/v1")
	{
		//注册
		v1.POST("/signup",controller.SignUpHandler)
		//登陆
		v1.POST("/login",controller.LoginHandler)
		//看帖子
		v1.GET("/post/:id",controller.GetPostHandler)
		//获取全部社区
		v1.GET("/classes",controller.ClassHandler)
		//看评论
		v1.GET("/post/:id/get_comment",controller.GetComment)
		//获取帖子列表
		v1.GET("/posts",controller.GetPostListHandler)
		//按照 时间 分数 获取帖子列表
		v1.GET("/posts_by",controller.GetPostListBy_Handler)
		//按照class获取帖子列表
		v1.GET("/posts_by_class",controller.GetPostListByClassIDHandler)
	}

	api:=r.Group("/blog/api")
	api.Use(middlewares.JWTAuthMiddleware())
	{
		//创建社区
		api.POST("/create_class",controller.CreateClass)
		//发帖子
		api.POST("/post",controller.CreatePostHandler)
		//发评论
		api.POST("/post/:id/add_comment",controller.AddCommentHandler)
		//投票
		api.POST("/vote",controller.VoteHandler)
	}

	return r
}
