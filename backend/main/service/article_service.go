package service

import (
	"context"
	"fiber-app/dto"
)

//For Finding Article
type IArticleService interface {
	GetArticle(ctx context.Context, topicID uint64) (dto.ArticleResponse, error)
}