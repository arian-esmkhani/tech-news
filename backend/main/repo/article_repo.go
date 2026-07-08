package repo

import (
	"context"
	"fiber-app/dto"
)

//For Finding Article
type IArticleRepo interface {
	GetArticle(ctx context.Context, topicID uint64) (dto.ArticleResponse, error)
}