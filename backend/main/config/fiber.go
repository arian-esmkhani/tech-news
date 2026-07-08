package config

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

var App *fiber.App

func SetupFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       		  true,
		CaseSensitive: 		  true,
		StrictRouting: 		  true,
		JSONEncoder:   		  sonic.Marshal,
		JSONDecoder:   	  	  sonic.Unmarshal,
		Network:       		  "tcp4",
		BodyLimit:             1 * 1024 * 1024,//1mb
		Concurrency:           70,
		DisableStartupMessage: false,
		ReduceMemoryUsage:     true,
		Immutable:             false,
		EnablePrintRoutes:     true,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
		AppName:               "PcPolojist",
		StreamRequestBody:     false,
	})
	
	App = app
	return app
}