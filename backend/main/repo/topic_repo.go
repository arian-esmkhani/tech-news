package repo

import (
	"context"
	"fiber-app/dto"
)

type ITopicRepo interface {
	//This method is for getting with category
	GetByTopic(ctx context.Context, name string, limit ,
	 offset  int64) ([]dto.DataDto, error)
	CountGetByTopic(ctx context.Context, name string) (int64, error)
	//This method is for searching with title or category
	SearchByTopic(ctx context.Context, name string, limit,
	 offset int64) ([]dto.DataDto, error)
	CountSearchByTopic(ctx context.Context, name string) (int64, error)
	//For Newest page with show the newest topic
	GetNewest(ctx context.Context, limit,
	 offset int64) ([]dto.DataDto, error)
	CountNewest(ctx context.Context) (int64, error)
}