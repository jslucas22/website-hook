package webserver

import (
	"website-conector/app/service"
	logging "website-conector/infra/logger"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type HttpServer struct {
	app         *fiber.App
	siteService *service.SiteService
	port        string
}

func NewHttpServer(
	port string,
	siteService *service.SiteService,
) *HttpServer {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		CaseSensitive:         true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	})
	app.Use(cors.New())
	return &HttpServer{
		app:         app,
		port:        port,
		siteService: siteService}
}

func (h *HttpServer) SetupRouter() {
	h.app.Get("/health", h.healthCheck)

	h.app.Get("/api/v1/sites", h.getSites)
	h.app.Get("/api/v1/sitesWithLogin", h.getSitesWithLogin)

	if err := h.app.Listen(":" + h.port); err != nil {
		logging.Fatal("setup-router", err.Error())
	}
}

func (h *HttpServer) healthCheck(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"status": "online"})
}
