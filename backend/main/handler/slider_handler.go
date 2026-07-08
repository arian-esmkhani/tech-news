package handler

import (
	"context"
	"fiber-app/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type SliderHandler struct {
	service service.ISliderService
}

func NewSliderHandler(service service.ISliderService) *SliderHandler{
	return &SliderHandler{service: service}
}

func (h *SliderHandler) NewestSlider(c *fiber.Ctx) error {

	reqCtx := c.Context()

	ctx, cancel := context.WithTimeout(reqCtx, 3*time.Second)
	defer cancel()

	slider, _ := h.service.NewestSlider(ctx)

	return c.Status(fiber.StatusOK).JSON(slider)
}

func (h *SliderHandler) OperatingSystemSlider(c *fiber.Ctx) error {

	reqCtx := c.Context()

	ctx, cancel := context.WithTimeout(reqCtx, 3*time.Second)
	defer cancel()

	slider, _ := h.service.OperatingSystemSlider(ctx)

	return c.Status(fiber.StatusOK).JSON(slider)
}

func (h *SliderHandler) LanguageSlider(c *fiber.Ctx) error {

	reqCtx := c.Context()

	ctx, cancel := context.WithTimeout(reqCtx, 3*time.Second)
	defer cancel()

	slider, _ := h.service.LanguageSlider(ctx)

	return c.Status(fiber.StatusOK).JSON(slider)
}

func (h *SliderHandler) JobSlider(c *fiber.Ctx) error {

	reqCtx := c.Context()

	ctx, cancel := context.WithTimeout(reqCtx, 3*time.Second)
	defer cancel()

	slider, _ := h.service.JobSlider(ctx)

	return c.Status(fiber.StatusOK).JSON(slider)
}

func (h *SliderHandler) AiSlider(c *fiber.Ctx) error {

	reqCtx := c.Context()

	ctx, cancel := context.WithTimeout(reqCtx, 3*time.Second)
	defer cancel()

	slider, _ := h.service.AiSlider(ctx)

	return c.Status(fiber.StatusOK).JSON(slider)
}