package route

import (
	"net/http"

	"github.com/an-halim/api-generate-report/handler"
	"github.com/an-halim/api-generate-report/repository"
	"github.com/an-halim/api-generate-report/service"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoute(app *fiber.App) {

	weatherRepo := repository.NewMeteorRemoteRepository(&http.Client{})
	weatherService := service.NewWeatherService(weatherRepo)
	weatherHandler := handler.NewWeatherHandler(weatherService)

	APIEXPORT := app.Group("/api/v1")
	APIEXPORT.Get("/report/pdf", weatherHandler.ReportPDF)
	APIEXPORT.Get("/report/csv", weatherHandler.ReportCSV)
}
