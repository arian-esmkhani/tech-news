package service

import (
	"context"
	"fiber-app/dto"
)

type ISliderService interface {
	NewestSlider(ctx context.Context) ([]dto.DataDto, error)
	OperatingSystemSlider(ctx context.Context) ([]dto.DataDto, error)
	LanguageSlider(ctx context.Context) ([]dto.DataDto, error)
	JobSlider(ctx context.Context) ([]dto.DataDto, error)
	AiSlider(ctx context.Context) ([]dto.DataDto, error)
}