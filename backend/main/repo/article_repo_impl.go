package repo

import (
	"context"
	"fiber-app/dto"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type articleRepo struct {
	db *sqlx.DB
}

func NewArticleRepo(db *sqlx.DB) IArticleRepo {
	return &articleRepo{db: db}
}

func (r *articleRepo) GetArticle(ctx context.Context, topicID uint64) (dto.ArticleResponse, error) {
	
	query := `
		SELECT
			t.title,
			t.description,
			t.imag_url AS topic_imag_url,
			t.first_category,
			t.second_category,
			a.head,
			a.body,
			a.imag_url AS article_imag_url,
			a.sec
		FROM topic t
		JOIN article a ON a.topic_id = t.id
		WHERE t.id = ?
		AND a.deleted_at IS null
		ORDER BY a.sec
    `

	var rows []dto.ArticleRow

	err := r.db.SelectContext(ctx, &rows, query, topicID)

	if len(rows) == 0 {
		return dto.ArticleResponse{}, fmt.Errorf("article not found for topicID %d", topicID)
	}

	first := rows[0]

	topic := dto.TopicDto{
		Title:          first.Title,
		Description:    first.Description,
		ImgUrl:         *first.TopicImage,
		FirstCategory:  first.FirstCategory,
		SecondCategory: first.SecondCategory,
	}

	articles := make([]dto.ArticleDto, 0, len(rows))
	for _, r := range rows {
		articles = append(articles, dto.ArticleDto{
			Head:     r.Head,
			Body:     r.Body,
			ImgUrl:   *r.ArticleImage,
			Sec:      r.Sec,
		})
	}

	return dto.ArticleResponse{
		TopicDto:    topic,
		ArticleDto:  articles,
	}, err
}