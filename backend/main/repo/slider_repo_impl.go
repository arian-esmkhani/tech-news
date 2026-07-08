package repo

import (
	"context"
	"fiber-app/dto"

	"github.com/jmoiron/sqlx"
)

type sliderRepo struct {
	db *sqlx.DB
}

func NewSliderRepo(db *sqlx.DB) ISliderRepo {
	return &sliderRepo{db: db}
}

const( 
	OPERATING_SYSTEM = "OperatingSystem";
    PROGRAMMING_LANGUAGE = "ProgrammingLanguage";
    JOB_POSITIONS = "JobPositions";
    AI = "Ai";
)

func (r sliderRepo) NewestSlider(ctx context.Context) ([]dto.DataDto, error) {

	query := `
        SELECT t.id, t.title, t.imag_url
        FROM topic t
        WHERE t.deleted_at IS NULL
        ORDER BY t.created_at DESC
		LIMIT 10
    `

	var slider []dto.DataDto

	err := r.db.SelectContext(ctx, &slider, query)

	return slider, err
}

func (r sliderRepo) OperatingSystemSlider(ctx context.Context) ([]dto.DataDto, error) {

	query := `
        SELECT t.id, t.title, t.imag_url
        FROM topic t
        WHERE t.deleted_at IS NULL AND t.first_category = ?
        ORDER BY t.created_at DESC
		LIMIT 10
    `

	var slider []dto.DataDto

	err := r.db.SelectContext(ctx, &slider, query, OPERATING_SYSTEM)

	return slider, err
}

func (r sliderRepo) LanguageSlider(ctx context.Context) ([]dto.DataDto, error) {

	query := `
        SELECT t.id, t.title, t.imag_url
        FROM topic t
        WHERE t.deleted_at IS NULL AND t.first_category = ?
        ORDER BY t.created_at DESC
		LIMIT 10
    `

	var slider []dto.DataDto

	err := r.db.SelectContext(ctx, &slider, query, PROGRAMMING_LANGUAGE)

	return slider, err
}

func (r sliderRepo) JobSlider(ctx context.Context) ([]dto.DataDto, error) {

	query := `
        SELECT t.id, t.title, t.imag_url
        FROM topic t
        WHERE t.deleted_at IS NULL AND t.first_category = ?
        ORDER BY t.created_at DESC
		LIMIT 10
    `

	var slider []dto.DataDto

	err := r.db.SelectContext(ctx, &slider, query, JOB_POSITIONS)

	return slider, err
}

func (r sliderRepo) AiSlider(ctx context.Context) ([]dto.DataDto, error) {

	query := `
        SELECT t.id, t.title, t.imag_url
        FROM topic t
        WHERE t.deleted_at IS NULL AND t.first_category = ?
        ORDER BY t.created_at DESC
		LIMIT 10
    `

	var slider []dto.DataDto

	err := r.db.SelectContext(ctx, &slider, query, AI)

	return slider, err
}