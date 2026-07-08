package service

import (
	"context"
	"fiber-app/dto"
)

type ITopicService interface {
	GetByTopic(ctx context.Context, name string, page int64) (dto.TopicResponse, error)
	SearchByTopic(ctx context.Context, name string, page int64) (dto.TopicResponse, error)
	GetNewest(ctx context.Context,page int64) (dto.TopicResponse, error)
}