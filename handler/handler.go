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

// ReportPDF godoc
//
//	@Summary		Generate a weather report in PDF format
//	@Description	Gets weather report data based on longitude and latitude and returns a downloadable PDF file.
//	@Tags			weather
//	@Accept			json
//	@Produce		application/pdf
//	@Param			long	query		string	true	"Longitude"
//	@Param			lat		query		string	true	"Latitude"
//	@Success		200		{file}		application/pdf
//	@Failure		400		{string}	map[string]interface{}	"Invalid input, longitude or latitude missing"
//	@Failure		500		{string}	map[string]interface{}	"Internal server error"
//	@Router			/report/pdf [get]
func (h *WeatherHandler) ReportPDF(c *fiber.Ctx) error {

	long := c.Query("long")
	lat := c.Query("lat")

	data, err := h.weatherService.ReportPDF(c.Context(), long, lat)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Download(data)
}

// ReportCSV godoc
//
//	@Summary		Generate a weather report in CSV format
//	@Description	Gets weather report data based on longitude and latitude and returns a downloadable CSV file.
//	@Tags			weather
//	@Accept			json
//	@Produce		text/csv
//	@Param			long	query		string	true	"Longitude"
//	@Param			lat		query		string	true	"Latitude"
//	@Success		200		{file}		text/csv
//	@Failure		400		{string}	map[string]interface{}	"Invalid input, longitude or latitude missing"
//	@Failure		500		{string}	map[string]interface{}	"Internal server error"
//	@Router			/report/csv [get]
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
