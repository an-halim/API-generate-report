package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"

	"github.com/an-halim/api-generate-report/repository"
)

type IWeatherService interface {
	ReportPDF(ctx context.Context, long string, lat string) (filePath string, err error)
	ReportCSV(ctx context.Context, long string, lat string) (filePath string, err error)
}

type WeatherService struct {
	weatherRemoteRepo repository.IRemoteRepository
	logger            *slog.Logger
}

func NewWeatherService(
	weatherRemoteRepo repository.IRemoteRepository,
) *WeatherService {
	return &WeatherService{
		weatherRemoteRepo: weatherRemoteRepo,
	}
}

func (s *WeatherService) ReportCSV(ctx context.Context, long string, lat string) (filePath string, err error) {

	weatherData, err := s.weatherRemoteRepo.Fetch(ctx, long, lat)

	fmt.Println(weatherData)
	if err != nil {

		return filePath, err
	}

	reportFile, err := os.Create("./weather_data.csv")
	if err != nil {
		return
	}
	defer reportFile.Close()

	fmt.Println(reportFile)
	csvWriter := csv.NewWriter(reportFile)
	defer csvWriter.Flush()

	// Write CSV headers
	headers := []string{"Time", "Temperature2M", "RelativeHumidity2M", "WindSpeed10M"}
	if err = csvWriter.Write(headers); err != nil {
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
			// slog.WarnContext(ctx, "error when writing CSV record", slog.Any("error", err))
			return filePath, err
		}
	}

	return reportFile.Name(), nil
}

func (s *WeatherService) ReportPDF(ctx context.Context, long string, lat string) (filePath string, err error) {

	return "", nil
}
