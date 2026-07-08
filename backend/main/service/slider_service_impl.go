package service

import (
	"context"
	"fiber-app/dto"
	"fiber-app/repo"

	"github.com/gofiber/fiber/v2/log"
)

type sliderService struct {
	slider repo.ISliderRepo
}

func NewSliderService(slider repo.ISliderRepo) ISliderService {
	return &sliderService{slider: slider}
}

func (s *sliderService) NewestSlider(ctx context.Context) ([]dto.DataDto, error) {

	slider, err := s.slider.NewestSlider(ctx)

    if err != nil {
		log.Errorf("Can not get the Newest Slider with error %v :", err)
        return nil, err
    }

	return slider, nil
}

func (s *sliderService) OperatingSystemSlider(ctx context.Context) ([]dto.DataDto, error) {

	slider, err := s.slider.OperatingSystemSlider(ctx)

    if err != nil {
		log.Errorf("Can not get the Operating System Slider with error %v :", err)
        return nil, err
    }

	return slider, nil
}

func (s *sliderService) LanguageSlider(ctx context.Context) ([]dto.DataDto, error) {

	slider, err := s.slider.LanguageSlider(ctx)

    if err != nil {
		log.Errorf("Can not get the Language Slider with error %v :", err)
        return nil, err
    }

	return slider, nil
}

func (s *sliderService) JobSlider(ctx context.Context) ([]dto.DataDto, error) {

	slider, err := s.slider.JobSlider(ctx)

    if err != nil {
		log.Errorf("Can not get the Job Slider with error %v :", err)
        return nil, err
    }

	return slider, nil
}

func (s *sliderService) AiSlider(ctx context.Context) ([]dto.DataDto, error) {

	slider, err := s.slider.AiSlider(ctx)

    if err != nil {
		log.Errorf("Can not get the Ai Slider with error %v :", err)
        return nil, err
    }

	return slider, nil
}