package service

import (
	"context"
	"fiber-app/dto"
	"fiber-app/repo"

	"github.com/gofiber/fiber/v2/log"
)

type articleService struct {
	article repo.IArticleRepo
}

func NewArticleService(article repo.IArticleRepo) IArticleService {
	return &articleService{article: article}
}

func (s *articleService) GetArticle(ctx context.Context, topicID uint64) (dto.ArticleResponse, error) {
	
	article, err := s.article.GetArticle(ctx, topicID)

	if err != nil {
		log.Errorf("Can not get the Article with error %v :", err)
        return dto.ArticleResponse{}, err
	}

	return article, nil
}

