package main

import (
	"log/slog"
	"net/http"

	"github.com/an-halim/api-generate-report/route"
	"github.com/gofiber/fiber/v2"
)

var (
	httpClient *http.Client
	logger     *slog.Logger
)

func main() {

	app := fiber.New()

	route.InitializeRoute(app)

	app.Listen(":3001")
}
