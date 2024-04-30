package webserver

import (
	"website-conector/infra/errors"

	"github.com/gofiber/fiber/v2"
)

func (h *HttpServer) getSites(c *fiber.Ctx) error {
	sites, err := h.siteService.GetSites()
	if err != nil {
		cError := errors.NewHttpError("gs-01: ", err.Error())
		return c.Status(400).JSON(cError)
	}
	return c.Status(200).JSON(sites)
}

func (h *HttpServer) getSitesWithLogin(c *fiber.Ctx) error {
	sites, err := h.siteService.GetSitesWithLogin()
	if err != nil {
		cError := errors.NewHttpError("gs-01: ", err.Error())
		return c.Status(400).JSON(cError)
	}
	return c.Status(200).JSON(sites)
}
