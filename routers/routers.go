package routers

import (
	"community_demo/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	//创建路由
	r := gin.Default()

	//路由组,处理请求
	v1 := r.Group("api/v1")
	{
		// 发布帖子
		v1.POST("topic", api.CreateTopic)          //发布话题
		v1.POST("topic/:tid/post", api.CreatePost) //回帖
		// 删除帖子
		v1.DELETE("topic/:tid", api.DeleteTopic)          //删除话题
		v1.DELETE("topic/:tid/post/:pid", api.DeletePost) //删除回复
		// 帖子话题运行修改
		v1.PUT("topic/:tid", api.UpdateTopic)
		v1.PUT("topic/:tid/post/:pid", api.UpdatePost)
		//展示话题和回帖列表
		v1.GET("topic/:tid", api.ShowTopic)
		v1.GET("topic/:tid/posts", api.ShowPost)

	}
	return r
}
