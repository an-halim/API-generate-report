package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/an-halim/api-generate-report/entity"
)

// IRemoteRepository mendefinisikan interface untuk repository pengguna
type IRemoteRepository interface {
	Fetch(ctx context.Context, long string, lat string) (res entity.WeatherData, err error)
}

// MeteorRemoteRepository mendefinisikan repository untuk mengambil data cuaca dari API
type MeteorRemoteRepository struct {
	httpClient *http.Client
}

// NewMeteorRemoteRepository membuat instance MeteorRemoteRepository
func NewMeteorRemoteRepository(httpClient *http.Client) *MeteorRemoteRepository {
	return &MeteorRemoteRepository{httpClient}
}

// Fetch mengambil data cuaca dari API
func (r *MeteorRemoteRepository) Fetch(ctx context.Context, long string, lat string) (res entity.WeatherData, err error) {

	if long == "" || lat == "" {
		err = errors.New("long or lat is empty")
		return res, err
	}

	endpoint := "https://api.open-meteo.com/v1/forecast?latitude=" + long + "&longitude=" + lat + "&current=temperature_2m,wind_speed_10m&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m"
	httpReq, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		fmt.Println(err)
		// logger.WarnContext(ctx, "error when hit http.NewRequestWithContext", err.Error())
		return res, err
	}

	data, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		fmt.Println("error when hit httpClient.Do", err.Error())
		// logger.WarnContext(ctx, "error when hit httpClient.Do", err.Error())
		return res, err
	}
	defer func() { _ = data.Body.Close() }()

	if data.StatusCode != http.StatusOK {
		errStatusCode := errors.New("not receiving status OK when hit API")
		// logger.WarnContext(ctx, errStatusCode.Error(), "err", err.Error())
		return res, errStatusCode
	}

	responseData, err := ioutil.ReadAll(data.Body)

	fmt.Println(responseData)
	if err != nil {
		fmt.Println("error when hit ioutil.ReadAll", err.Error())
		// logger.WarnContext(ctx, "error when hit ioutil.ReadAll", err.Error())
		return res, err
	}
	var weather entity.WeatherData
	if err = json.Unmarshal(responseData, &weather); err != nil {
		log.Println("error when hit json.Unmarshal", err.Error())

		return res, err
	}

	fmt.Println(weather)

	return weather, nil
}
