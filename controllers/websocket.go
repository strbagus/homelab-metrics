package controllers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/strbagus/homelab-metrics/utils"
)

func WSTop(c *websocket.Conn) {
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
}
