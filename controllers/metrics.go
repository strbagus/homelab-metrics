package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/strbagus/homelab-metrics/utils"
)

func GetTop(c *fiber.Ctx) error {
	metrics := utils.GetMetric()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"title":    "Homelab Metric",
		"data":     metrics,
		"datetime": time.Now().UTC().Format(time.RFC3339),
	})
}
func GetNodes(c *fiber.Ctx) error {
	nodes := utils.GetNodes()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"title":    "Homelab Nodes",
		"data":     nodes,
		"datetime": time.Now().UTC().Format(time.RFC3339),
	})
}
func GetPods(c *fiber.Ctx) error {
	pods := utils.GetPods()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"title":    "Homelab Pods",
		"data":     pods,
		"datetime": time.Now().UTC().Format(time.RFC3339),
	})
}
func GetPodKinds(c *fiber.Ctx) error {
	kinds := utils.GetPodKinds()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"title":    "Homelab Kinds",
		"data":     kinds,
		"datetime": time.Now().UTC().Format(time.RFC3339),
	})
}
func GetServices(c *fiber.Ctx) error {
	services := utils.GetInfoServices()
	host := utils.GetHost()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"title":    "Homelab Systemd Services",
		"data":     services,
		"host":     host,
		"datetime": time.Now().UTC().Format(time.RFC3339),
	})
}
