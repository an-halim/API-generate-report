package service

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/an-halim/api-generate-report/repository"
	"github.com/gofiber/fiber/v2/log"
)

type IWeatherService interface {
	ReportPDF(ctx context.Context, long string, lat string) (filePath string, err error)
	ReportCSV(ctx context.Context, long string, lat string) (filePath string, err error)
}

type WeatherService struct {
	weatherRemoteRepo repository.IRemoteRepository
}

func NewWeatherService(
	weatherRemoteRepo repository.IRemoteRepository,
) *WeatherService {
	return &WeatherService{
		weatherRemoteRepo: weatherRemoteRepo,
	}
}

func (s *WeatherService) ReportCSV(ctx context.Context, long string, lat string) (filePath string, err error) {

	log.Info("Generating CSV report...")
	weatherData, err := s.weatherRemoteRepo.Fetch(ctx, long, lat)

	if err != nil {
		log.Error(err)
		return filePath, err
	}

	// Create CSV file
	reportFile, err := os.Create(fmt.Sprintf("./report_%s.csv", weatherData.Hourly.Time[0]))
	if err != nil {
		log.Error(err)
		return
	}
	defer reportFile.Close()

	csvWriter := csv.NewWriter(reportFile)
	defer csvWriter.Flush()

	log.Info("Writing CSV report to file...")

	// Write CSV headers
	headers := []string{"Time", "Temperature2M", "RelativeHumidity2M", "WindSpeed10M"}
	if err = csvWriter.Write(headers); err != nil {
		log.Error(err)
		return filePath, err
	}

	// Write hourly data
	for i, time := range weatherData.Hourly.Time {
		record := []string{
			time,
			fmt.Sprintf("%.2f", weatherData.Hourly.Temperature2M[i]),
			fmt.Sprintf("%d", weatherData.Hourly.RelativeHumidity2M[i]),
			fmt.Sprintf("%.2f", weatherData.Hourly.WindSpeed10M[i]),
		}
		if err = csvWriter.Write(record); err != nil {
			log.Error(err)
			return filePath, err
		}
	}

	log.Info("CSV report generated successfully!")

	return reportFile.Name(), nil
}

func (s *WeatherService) ReportPDF(ctx context.Context, long string, lat string) (filePath string, err error) {

	log.Info("Generating PDF report...")

	data, err := s.weatherRemoteRepo.Fetch(ctx, long, lat)

	if err != nil {
		log.Error(err)
		return filePath, err
	}

	template, err := template.ParseFiles("./static/report.html")

	if err != nil {
		log.Error(err)
		return filePath, err
	}

	buf := new(bytes.Buffer)

	if err = template.Execute(buf, data); err != nil {
		log.Error(err)

		return filePath, err
	}

	pdfg, er := wkhtmltopdf.NewPDFGenerator()

	if er != nil {
		fmt.Println(er)
		return filePath, er
	}

	page := wkhtmltopdf.NewPageReader(strings.NewReader(buf.String()))

	pdfg.AddPage(page)

	if err = pdfg.Create(); err != nil {
		log.Error(err)
		return filePath, err
	}

	// generate time
	filePath = fmt.Sprintf("./report_%s.pdf", data.Hourly.Time[0])

	log.Info("Writing PDF report to file...")

	if err = pdfg.WriteFile(filePath); err != nil {
		log.Error(err)
		return filePath, err
	}

	log.Info("PDF report generated successfully!")

	return filePath, nil
}
