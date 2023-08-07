package api

import (
	"community_demo/service"
	"github.com/gin-gonic/gin"
	"log"
)

// 新增一个回复
func CreatePost(ctx *gin.Context) {
	var createPost service.CreatePostService
	//ShouldBind绑定到指定的结构体对象
	err := ctx.ShouldBind(&createPost)
	if err == nil { //绑定成功
		//新增帖子
		res := createPost.Create(ctx.Param("tid"))
		ctx.JSON(200, res)
	} else {
		log.Println(err)
		//返回错误信息
		ctx.JSON(400, ErrorResponse(err))
	}
}

// 删除一个回复
func DeletePost(ctx *gin.Context) {
	var deletePost service.DeletePostService
	err := ctx.ShouldBind(&deletePost)
	if err == nil { //绑定成功
		//删除回复
		res := deletePost.Delete(ctx.Param("tid"), ctx.Param("pid"))
		ctx.JSON(200, res)
	} else {
		log.Println(err)
		//返回错误信息
		ctx.JSON(400, ErrorResponse(err))
	}
}

// 修改一个回复
func UpdatePost(ctx *gin.Context) {
	var updatePost service.UpdatePostService
	err := ctx.ShouldBind(&updatePost)
	if err == nil { //绑定成功
		//修改回复
		res := updatePost.Update(ctx.Param("tid"), ctx.Param("pid"))
		ctx.JSON(200, res)
	} else {
		log.Println(err)
		//返回错误信息
		ctx.JSON(400, ErrorResponse(err))
	}
}

// 展示话题下所有回复
func ShowPost(ctx *gin.Context) {
	var showPost service.PostListService
	err := ctx.ShouldBind(&showPost)
	if err == nil { //绑定成功
		//展示回复
		res := showPost.Show(ctx.Param("tid"))
		ctx.JSON(200, res)
	} else {
		log.Println(err)
		//返回错误信息
		ctx.JSON(400, ErrorResponse(err))
	}
}
