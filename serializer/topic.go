package serializer

import "community_demo/models"

type Topic struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Status  int    `json:"status"`
	Content string `json:"content"`
}

func BuildTopic(item models.Topic) Topic {
	return Topic{
		ID:      item.ID,
		Title:   item.Title,
		Status:  item.Status,
		Content: item.Content,
	}
}

func BuildTopics(items []models.Topic) (tasks []Topic) {
	for _, item := range items {
		task := BuildTopic(item)
		tasks = append(tasks, task)
	}
	return tasks
}
