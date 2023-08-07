package service

import (
	"community_demo/models"
	"community_demo/serializer"
	"strconv"
)

// 创回复
type CreatePostService struct {
	TopicId int    `json:"topic_id" form:"topic_id"` //外键关联id
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` // 0是未做，1是已做
}

func (service *CreatePostService) Create(topicId string) serializer.Response {
	tid, err := strconv.Atoi(topicId)

	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "话题ID无效",
		}
	}

	var topic models.Topic
	err = models.DB.First(&topic, tid).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "找不到指定话题",
		}
	}

	post := models.Post{
		TopicId: tid,
		Content: service.Content,
		Status:  0,
	}
	//数据库操作 插入话题表单
	err = models.DB.Create(&post).Error

	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "回帖失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "评论成功",
	}
}

// 删回复
type DeletePostService struct {
}

func (service *DeletePostService) Delete(topicId string, positionStr string) serializer.Response {
	// 将字符串类型的 position 转换为 int 类型
	position, _ := strconv.Atoi(positionStr)
	//默认第一行为0
	tid, err := strconv.Atoi(topicId)

	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "话题ID无效",
		}
	}

	var posts []models.Post
	err = models.DB.Where("topic_id = ?", tid).Find(&posts).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "获取话题回复失败",
		}
	}

	if len(posts) == 0 {
		return serializer.Response{
			Status: 404,
			Msg:    "找不到指定话题",
		}
	}

	// 确保 position 不超出回复列表范围
	if position < 1 || position >= len(posts) {
		return serializer.Response{
			Status: 400,
			Msg:    "回复位置无效",
		}
	}

	// 删除指定位置的回复
	err = models.DB.Delete(&posts[position-1]).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "评论删除失败",
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "评论删除成功",
	}
}

// 改回复
type UpdatePostService struct {
	TopicId int    `json:"topic_id" form:"topic_id"` //外键关联id
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` // 0是未做，1是已做
}

func (service *UpdatePostService) Update(topicId string, positionStr string) serializer.Response {
	// 将字符串类型的 position 转换为 int 类型
	position, _ := strconv.Atoi(positionStr)

	tid, err := strconv.Atoi(topicId)
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "话题ID无效",
		}
	}

	var posts []models.Post
	err = models.DB.Where("topic_id = ?", tid).Find(&posts).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "获取话题回复失败",
		}
	}

	// 确保 position 不超出回复列表范围
	if position < 1 || position >= len(posts) {
		return serializer.Response{
			Status: 400,
			Msg:    "回复位置无效",
		}
	}

	// 获取要更新的回复
	// position 1 对应 posts 0
	post := posts[position-1]

	// 更新回复内容和状态
	post.Content = service.Content
	post.Status = service.Status

	// 保存更新后的回复
	err = models.DB.Save(&post).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "修改失败",
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "修改成功",
	}
}

type PostListService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

func (service *PostListService) Show(topicId string) serializer.Response {
	var posts []models.Post

	// 判断是否需要分页
	if service.PageSize == 0 {
		service.PageSize = -1
	}
	if service.PageNum == 0 {
		service.PageNum = -1
	}
	offsetVal := (service.PageNum - 1) * service.PageSize
	if service.PageNum == -1 && service.PageSize == -1 {
		offsetVal = -1
	}

	// 返回一个总数
	var total int64
	// 查询数据库
	tid, err := strconv.Atoi(topicId)
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "话题ID无效",
		}
	}
	query := models.DB.Where("topic_id = ?", tid)

	//
	if offsetVal >= 0 {
		query = query.Offset(offsetVal)
	}
	if service.PageSize >= 0 {
		query = query.Limit(service.PageSize)
	}
	// 查询并将结果赋值给posts
	query.Find(&posts).Count(&total)
	// 获取回复列表的总记录数
	models.DB.Model(&models.Post{}).Where("topic_id = ?", tid).Count(&total)

	return serializer.BuildListResponse(serializer.BuildPosts(posts), uint(total))
}
