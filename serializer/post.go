package serializer

import "community_demo/models"

type Post struct {
	ID      uint   `json:"id"`
	TopicId int    `json:"topic_id"`
	Status  int    `json:"status"`
	Content string `json:"content"`
}

func BuildPost(item models.Post) Post {
	return Post{
		ID:      item.ID,
		TopicId: item.TopicId,
		Status:  item.Status,
		Content: item.Content,
	}
}

func BuildPosts(items []models.Post) (posts []Post) {
	for _, item := range items {
		post := BuildPost(item)
		posts = append(posts, post)
	}
	return posts
}
