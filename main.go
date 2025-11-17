package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/strbagus/homelab-metrics/controllers"
	"github.com/strbagus/homelab-metrics/middlewares"
	"github.com/strbagus/homelab-metrics/utils"
)

func init() {
	if os.Getenv("ENV") != "production" {
		godotenv.Load()
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file: ", err)
		}
	}
}

type Client struct {
	conn *websocket.Conn
}

type Message struct {
	Name     string             `json:"name"`
	Datetime string             `json:"datetime"`
	Nodes    []utils.NodeMetric `json:"nodes"`
}

func main() {
	app := fiber.New(fiber.Config{})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CORS_ALLOWED_ORIGIN"),
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))

	ping := app.Group("/ping")
	ping.Post("/disk", controllers.PingDisk)
	base := app.Group(os.Getenv("BASE_URL"))
	base.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Halo Dunia!")
	})
	base.Get("/ws/top", websocket.New(controllers.WSTop))
	v1 := base.Group("/v1", middlewares.JWTMiddleware())
	v1.Get("/top", controllers.GetTop)
	v1.Get("/nodes", controllers.GetNodes)
	v1.Get("/nodes", controllers.GetNodes)
	v1.Get("/kinds", controllers.GetPodKinds)
	v1.Get("/pods", controllers.GetPods)
	v1.Post("/detail", controllers.GetDetail)
	v1.Get("/services", controllers.GetServices)
	v1.Get("/disks", controllers.GetDisks)

	listenPort := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))
	log.Fatal(app.Listen(listenPort))
}
