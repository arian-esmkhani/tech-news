package service

import (
	"context"
	"fiber-app/dto"
	"fiber-app/helper"
	"fiber-app/repo"

	"github.com/gofiber/fiber/v2/log"
)

type topicService struct {
	topic repo.ITopicRepo
}

func NewTopicService(topic repo.ITopicRepo) ITopicService {
	return &topicService{topic: topic}
}

const limitPage int64 = 12

func (t *topicService) GetByTopic(ctx context.Context, name string,
	page int64) (dto.TopicResponse, error) {

	count, err := t.topic.CountGetByTopic(ctx, name)
	var hasNext bool = helper.HasNext(page, count, limitPage)

	topic, err := t.topic.GetByTopic(ctx, name, limitPage, (page*limitPage))

	if err != nil {
		log.Errorf("Error in getting by topic with error %v :", err)
	}

	response := dto.TopicResponse { DataDto: topic, HasNext: hasNext }

	return response, nil
}

func (t *topicService) SearchByTopic(ctx context.Context, name string,
	page int64) (dto.TopicResponse, error) {

	count, err := t.topic.CountSearchByTopic(ctx, name)
	var hasNext bool = helper.HasNext(page, count, limitPage)

	topic, err := t.topic.SearchByTopic(ctx, name, limitPage, (page*limitPage))

	if err != nil {
		log.Errorf("Error in searching by topic with error %v :", err)
	}

	response := dto.TopicResponse { DataDto: topic, HasNext: hasNext }

	return response, nil
}

func (t *topicService) GetNewest(ctx context.Context,
	page int64) (dto.TopicResponse, error) {

	count, err := t.topic.CountNewest(ctx)
	var hasNext bool = helper.HasNext(page, count, limitPage)

	topic, err := t.topic.GetNewest(ctx, limitPage, (page*limitPage))

	if err != nil {
		log.Errorf("Error in getting by topic with error %v :", err)
	}

	response := dto.TopicResponse { DataDto: topic, HasNext: hasNext }

	return response, nil
}