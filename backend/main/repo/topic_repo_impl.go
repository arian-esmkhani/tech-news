package repo

import (
	"context"
	"fiber-app/dto"
	"strings"

	"github.com/jmoiron/sqlx"
)

type topicRepo struct {
	db *sqlx.DB
}

func NewTopicRepo(db *sqlx.DB) ITopicRepo {
	return &topicRepo{db: db}
}

func (r *topicRepo) GetByTopic(ctx context.Context, name string, limit,
	offset int64) ([]dto.DataDto, error) {

	query := `
        SELECT t.id, t.title, t.imag_url
		FROM topic t
		WHERE t.deleted_at IS NULL
		AND( lower(t.first_category) =  lower(?)
		or lower(t.second_category) =  lower(?) )
		ORDER BY t.id ASC
		LIMIT ? OFFSET ?
    `

	var topic []dto.DataDto

	err := r.db.SelectContext(ctx, &topic, query, name, name, limit, offset)

	return topic, err
}

func (r *topicRepo) CountGetByTopic(ctx context.Context, name string) (int64, error) {

	query := `
        SELECT count(*)
		FROM topic t
		WHERE t.deleted_at IS NULL
		AND( lower(t.first_category) =  lower(?)
		or lower(t.second_category) =  lower(?) )
    `

	var count int64

	err := r.db.SelectContext(ctx, &count, query, name, name)

	return count, err
}

func (r *topicRepo) SearchByTopic(ctx context.Context, name string, limit,
	offset int64) ([]dto.DataDto, error) {
	query := `
        SELECT t.id, t.title, t.imag_url
        FROM topic t
        WHERE t.deleted_at IS NULL
    `

	var params []any

	if name != "" {
		trimmedName := strings.TrimSpace(name)
		query += `
            AND (LOWER(t.title) LIKE LOWER(CONCAT('%', ?, '%'))
                 OR LOWER(t.first_category) LIKE LOWER(CONCAT('%', ?, '%'))
                 OR LOWER(t.second_category) LIKE LOWER(CONCAT('%', ?, '%')))`

		params = append(params, trimmedName, trimmedName, trimmedName)
	}

	query += " LIMIT ? OFFSET ?"
	params = append(params, limit, offset)

	var results []dto.DataDto

	err := r.db.SelectContext(ctx, &results, query, params...)

	return results, err
}

func (r *topicRepo) CountSearchByTopic(ctx context.Context, name string) (int64, error) {
	query := `
        SELECT count(*)
        FROM topic t
        WHERE t.deleted_at IS NULL
    `

	var params []any

	if name != "" {
		trimmedName := strings.TrimSpace(name)
		query += `
            AND (LOWER(t.title) LIKE LOWER(CONCAT('%', ?, '%'))
                 OR LOWER(t.first_category) LIKE LOWER(CONCAT('%', ?, '%'))
                 OR LOWER(t.second_category) LIKE LOWER(CONCAT('%', ?, '%')))`

		params = append(params, trimmedName, trimmedName, trimmedName)
	}

	var count int64

	err := r.db.SelectContext(ctx, &count, query, params...)

	return count, err
}

func (r *topicRepo) GetNewest(ctx context.Context, limit,
	offset int64) ([]dto.DataDto, error) {

	query := `
        SELECT t.id, t.title, t.imag_url
		FROM topic t
		WHERE t.deleted_at IS NULL
		ORDER BY t.created_at DESC
		LIMIT ? OFFSET ?
    `

	var topic []dto.DataDto

	err := r.db.SelectContext(ctx, &topic, query, limit, offset)

	return topic, err
}

func (r *topicRepo) CountNewest(ctx context.Context) (int64, error) {

	query := `
        SELECT count(*)
		FROM topic t
		WHERE t.deleted_at IS NULL
    `

	var count int64

	err := r.db.SelectContext(ctx, &count, query)

	return count, err
}