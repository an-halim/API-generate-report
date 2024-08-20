package entity

type WeatherData struct {
	Latitude             float64        `json:"latitude"`
	Longitude            float64        `json:"longitude"`
	GenerationTimeMs     float64        `json:"generationtime_ms"`
	UtcOffsetSeconds     int            `json:"utc_offset_seconds"`
	Timezone             string         `json:"timezone"`
	TimezoneAbbreviation string         `json:"timezone_abbreviation"`
	Elevation            float64        `json:"elevation"`
	CurrentUnits         Units          `json:"current_units"`
	Current              CurrentWeather `json:"current"`
	HourlyUnits          Units          `json:"hourly_units"`
	Hourly               HourlyForecast `json:"hourly"`
}

type Units struct {
	Time               string `json:"time"`
	Interval           string `json:"interval,omitempty"`
	Temperature2M      string `json:"temperature_2m"`
	WindSpeed10M       string `json:"wind_speed_10m,omitempty"`
	RelativeHumidity2M string `json:"relative_humidity_2m,omitempty"`
}

type CurrentWeather struct {
	Time          string  `json:"time"`
	Interval      int     `json:"interval"`
	Temperature2M float64 `json:"temperature_2m"`
	WindSpeed10M  float64 `json:"wind_speed_10m"`
}

type HourlyForecast struct {
	Time               []string  `json:"time"`
	Temperature2M      []float64 `json:"temperature_2m"`
	RelativeHumidity2M []int     `json:"relative_humidity_2m"`
	WindSpeed10M       []float64 `json:"wind_speed_10m"`
}
