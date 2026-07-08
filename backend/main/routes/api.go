package routes

import (
	"fiber-app/handler"
	"fiber-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupAPIRoutes(app *fiber.App,
	sliderHandler *handler.SliderHandler, topicHandler *handler.TopicHandler,
	articleHandler *handler.ArticleHandler) {

	api := app.Group("/api",
		middleware.Compression(),
		middleware.CORS(),
		middleware.Limiter(),
	)

	slider := api.Group("/slider")
	slider.Get("/newest", sliderHandler.NewestSlider)
	slider.Get("/system", sliderHandler.OperatingSystemSlider)
	slider.Get("/language", sliderHandler.LanguageSlider)
	slider.Get("/job", sliderHandler.JobSlider)
	slider.Get("/ai", sliderHandler.AiSlider)

	topic := api.Group("/topic")
	topic.Get("/by-topic/:name", topicHandler.GetByTopic)
	topic.Get("/search-topic/:name", topicHandler.SearchByTopic)
	topic.Get("/newest", topicHandler.GetNewest)

	article := api.Group("/article")
	article.Get("/get/:topicID", articleHandler.GetArticle)
}