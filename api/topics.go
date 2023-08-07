package api

import (
	"community_demo/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 新增一个帖子
func CreateTopic(ctx *gin.Context) {
	var createTopic service.CreateTopicService
	//ShouldBind绑定到指定的结构体对象
	err := ctx.ShouldBind(&createTopic)
	if err == nil { //绑定成功
		//新增帖子
		res := createTopic.Create()
		ctx.JSON(http.StatusOK, res)
	} else {
		log.Println(err)
		//返回错误信息
		ctx.JSON(400, ErrorResponse(err))
	}
}

// 删除一个帖子
func DeleteTopic(ctx *gin.Context) {
	var deleteTopic service.DeleteTopicService
	//ShouldBind绑定到指定的结构体对象
	err := ctx.ShouldBind(&deleteTopic)
	if err == nil { //绑定成功
		//新增帖子
		res := deleteTopic.Delete(ctx.Param("tid"))
		ctx.JSON(200, res)
	} else {
		log.Println(err)
		//返回错误信息
		ctx.JSON(400, ErrorResponse(err))
	}
}

// 帖子话题修改
func UpdateTopic(ctx *gin.Context) {
	var updateTopic service.UpdateTopicService
	//ShouldBind绑定到指定的结构体对象
	err := ctx.ShouldBind(&updateTopic)
	if err == nil { //绑定成功
		//新增帖子
		res := updateTopic.Update(ctx.Param("tid"))
		ctx.JSON(200, res)
	} else {
		log.Println(err)
		//返回错误信息
		ctx.JSON(400, ErrorResponse(err))
	}
}

// 展示一条话题
func ShowTopic(ctx *gin.Context) {
	var showTopic service.ShowTopicService
	err := ctx.ShouldBind(&showTopic)
	if err == nil {
		res := showTopic.Show(ctx.Param("tid"))
		ctx.JSON(200, res)
	} else {
		log.Println(err)
		//返回错误信息
		ctx.JSON(400, ErrorResponse(err))
	}
}
