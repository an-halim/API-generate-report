package handler

import (
	"github.com/an-halim/api-generate-report/service"
	"github.com/gofiber/fiber/v2"
)

type WeatherHandler struct {
	weatherService service.IWeatherService
}

func NewWeatherHandler(
	weatherService service.IWeatherService,
) *WeatherHandler {
	return &WeatherHandler{
		weatherService: weatherService,
	}
}

func (h *WeatherHandler) ReportPDF(c *fiber.Ctx) error {

	long := c.Query("long")
	lat := c.Query("lat")

	data, err := h.weatherService.ReportPDF(c.Context(), long, lat)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Download(data, "report.pdf")
}

func (h *WeatherHandler) ReportCSV(c *fiber.Ctx) error {
	long := c.Query("long")
	lat := c.Query("lat")

	if long == "" || lat == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "long or lat is empty",
		})
	}

	filePath, err := h.weatherService.ReportCSV(c.Context(), long, lat)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Download(filePath)
}
