package handler

import (
	"context"
	"fiber-app/service"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type TopicHandler struct {
	service service.ITopicService
}

func NewTopicHandler(service service.ITopicService) *TopicHandler{
	return &TopicHandler{service: service}
}

func (h *TopicHandler) GetByTopic(c *fiber.Ctx) error {

	reqCtx := c.Context()

	name := c.Params("name")

	if name == "" || strings.TrimSpace(name) == ""  {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   "Name is null aor blank",
		})
	}

	page, err := strconv.Atoi(c.Query("page", "0"))

	if err != nil || page < 1 {
		page = 0
	}

	if page > 1000 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   "Invalid page number",
			"Message": "Page number is too large",
		})
	}

	ctx, cancel := context.WithTimeout(reqCtx, 3*time.Second)
	defer cancel()

	topic, _ := h.service.GetByTopic(ctx, name, int64(page))

	if len(topic.DataDto) <= 0 {
		return c.Status(fiber.StatusNotFound).JSON(topic)
	}

	return c.Status(fiber.StatusOK).JSON(topic)
}

func (h *TopicHandler) SearchByTopic(c *fiber.Ctx) error {

	reqCtx := c.Context()

	name := c.Params("name")

	if name == "" || strings.TrimSpace(name) == ""  {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   "Name is null aor blank",
		})
	}

	page, err := strconv.Atoi(c.Query("page", "0"))

	if err != nil || page < 1 {
		page = 0
	}

	if page > 1000 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   "Invalid page number",
			"Message": "Page number is too large",
		})
	}

	ctx, cancel := context.WithTimeout(reqCtx, 3*time.Second)
	defer cancel()

	topic, _ := h.service.SearchByTopic(ctx, name, int64(page))

	if len(topic.DataDto) <= 0 {
		return c.Status(fiber.StatusNoContent).JSON(topic)
	}

	return c.Status(fiber.StatusOK).JSON(topic)
}

func (h *TopicHandler) GetNewest(c *fiber.Ctx) error {

	reqCtx := c.Context()

	page, err := strconv.Atoi(c.Query("page", "0"))

	if err != nil || page < 1 {
		page = 0
	}

	if page > 1000 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   "Invalid page number",
			"Message": "Page number is too large",
		})
	}

	ctx, cancel := context.WithTimeout(reqCtx, 3*time.Second)
	defer cancel()

	topic, _ := h.service.GetNewest(ctx, int64(page))

	if len(topic.DataDto) <= 0 {
		return c.Status(fiber.StatusNotFound).JSON(topic)
	}

	return c.Status(fiber.StatusOK).JSON(topic)
}