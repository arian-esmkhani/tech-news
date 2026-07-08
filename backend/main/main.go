package main

import (
	"fiber-app/config"
	"fiber-app/handler"
	"fiber-app/repo"
	"fiber-app/routes"
	"fiber-app/service"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	
	db, err := config.DbSetup()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer db.Close()

	config.SetupSonic()
	app   := config.SetupFiber()

	sliderRepo 	   := repo.NewSliderRepo(db)
	topicRepo 	   := repo.NewTopicRepo(db)
	articleRepo	   := repo.NewArticleRepo(db)
	sliderService    := service.NewSliderService(sliderRepo)
	topicService    := service.NewTopicService(topicRepo)
	articleService  := service.NewArticleService(articleRepo)
	sliderHandler    := handler.NewSliderHandler(sliderService)
	topicHandler    := handler.NewTopicHandler(topicService)
	articleHandler  := handler.NewArticleHandler(articleService)

	routes.SetupAPIRoutes(app, sliderHandler, topicHandler, articleHandler)

	log.Fatal(app.Listen(":5002"))
}
