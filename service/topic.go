package service

import (
	"community_demo/models"
	"community_demo/serializer"
	"fmt"
)

type CreateTopicService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` // 0是未做，1是已做
}

func (service *CreateTopicService) Create() serializer.Response {

	topic := models.Topic{
		Title:   service.Title,
		Content: service.Content,
		Status:  0,
	}
	fmt.Println("准备创建topic")

	sqlDb, err := models.DB.DB()
	if err != nil {
		fmt.Println("获取数据库连接池失败：", err)
		return serializer.Response{
			Status: 500,
			Msg:    "获取数据库连接池失败",
		}
	}
	fmt.Println("连接池状态：", sqlDb.Stats())

	err = models.DB.Create(&topic).Error

	if err != nil {
		fmt.Println("创建topic失败：", err)
		return serializer.Response{
			Status: 500,
			Msg:    "创建话题失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建成功",
		Data:   topic, // 设置"data"字段为创建的话题信息
	}
}

type DeleteTopicService struct {
}

// 删除一个帖子
func (service *DeleteTopicService) Delete(tid string) serializer.Response {
	var topic models.Topic

	//数据库表单中删除tid帖子
	err := models.DB.Delete(&topic, tid).Error

	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}

type UpdateTopicService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` // 0是未做，1是已做
}

// 更新一条帖子
func (service *UpdateTopicService) Update(tid string) serializer.Response {
	var topic models.Topic
	models.DB.First(&topic, tid)
	topic.Content = service.Content
	topic.Title = service.Title
	topic.Status = service.Status
	err := models.DB.Save(&topic).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据保存出错",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTopic(topic),
		Msg:    "更新完成",
	}
}

// 展示一个帖子话题
type ShowTopicService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` // 0是未做，1是已做
}

func (service *ShowTopicService) Show(tid string) serializer.Response {
	var topic models.Topic
	//找到话题
	err := models.DB.First(&topic, tid).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "查询失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTopic(topic),
	}
}
