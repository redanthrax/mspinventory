package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleClients(c *fiber.Ctx) error {
  return c.Render("clients/list", fiber.Map{}, "layouts/main")
}

func HandleClientsAdd(c *fiber.Ctx) error {
  return c.Render("clients/add", fiber.Map{}, "layouts/main")
}
