package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	m "github.com/strbagus/homelab-metrics/models"
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

type RequestDetail struct {
	Name string `json:"name"`
}

func GetDetail(c *fiber.Ctx) error {
	var data RequestDetail
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err,
		})
	}
	detail := utils.GetDetail(data.Name)
	name := fmt.Sprintf("Detail %v", data.Name)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"title": name,
		"data":  detail,
	})
}

func PingDisk(c *fiber.Ctx) error {
	var data m.DiskType
	err := c.BodyParser(&data)

	res := utils.AddDiskPing(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"err":     err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}
