package handler

import (
	"context"
	"fiber-app/service"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ArticleHandler struct {
	service service.IArticleService
}

func NewArticleHandler(service service.IArticleService) *ArticleHandler {
	return &ArticleHandler{service: service}
}

func (h *ArticleHandler) GetArticle(c *fiber.Ctx) error {
	
	reqCtx := c.Context()

	topicID, _ := strconv.Atoi(c.Params("topicID"))

	if topicID <= 0  || topicID > 100000 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error": "er"})
	}

	ctx, cancel := context.WithTimeout(reqCtx, 3*time.Second)
	defer cancel()

	article, _ := h.service.GetArticle(ctx, uint64(topicID))

	if len(article.ArticleDto) == 0{
		return c.Status(fiber.StatusNoContent).JSON(article)
	}

	return c.Status(fiber.StatusOK).JSON(article)
}