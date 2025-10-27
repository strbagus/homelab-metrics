package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
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
		AllowMethods:     "GET",
		AllowCredentials: true,
	}))
	base := app.Group(os.Getenv("BASE_URL"))
	v1 := base.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Halo Dunia!")
	})
	v1.Get("/top", func(c *fiber.Ctx) error {
		metrics := utils.GetMetric()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"title":    "Homelab Metric",
			"data":     metrics,
			"datetime": time.Now().UTC().Format(time.RFC3339),
		})
	})
	v1.Get("/nodes", func(c *fiber.Ctx) error {
		nodes := utils.GetNodes()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"title":    "Homelab Nodes",
			"data":     nodes,
			"datetime": time.Now().UTC().Format(time.RFC3339),
		})
	})
	v1.Get("/pods", func(c *fiber.Ctx) error {
		pods := utils.GetPods()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"title":    "Homelab Pods",
			"data":     pods,
			"datetime": time.Now().UTC().Format(time.RFC3339),
		})
	})
	v1.Get("/kinds", func(c *fiber.Ctx) error {
		kinds := utils.GetPodKinds()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"title":    "Homelab Kinds",
			"data":     kinds,
			"datetime": time.Now().UTC().Format(time.RFC3339),
		})
	})
	v1.Get("/ws/metrics", websocket.New(func(c *websocket.Conn) {
		defer c.Close()
		log.Println("[INFO] Client connected to metrics ws")

		if data, err := json.Marshal(utils.GetMetric()); err == nil {
			c.WriteMessage(websocket.TextMessage, data)
		}

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		done := make(chan struct{})

		go func() {
			for {
				select {
				case <-ticker.C:
					data, _ := json.Marshal(utils.GetMetric())
					if err := c.WriteMessage(websocket.TextMessage, data); err != nil {
						log.Println("[ERROR] write:", err)
						close(done)
						return
					}
				case <-done:
					return
				}
			}
		}()

		for {
			if _, _, err := c.ReadMessage(); err != nil {
				log.Println("[INFO] Client disconnected")
				close(done)
				break
			}
		}
	}))
	listenPort := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))
	log.Fatal(app.Listen(listenPort))
}
