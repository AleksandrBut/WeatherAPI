package client

import (
	"WeatherAPI/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

const externalApiBaseUrl = "https://api.weatherapi.com/v1"
const currentWeatherRequestUrl = "/current.json"
const queryParam = "q"
const apiKeyParam = "key"
const airQualityParam = "aqi"
const no = "no"

var apiKey = os.Getenv("EXTERNAL_WEATHER_API_KEY")

func GetWeatherByCity(cityName string) (model.Weather, error) {
	request, err := http.NewRequest(http.MethodGet, externalApiBaseUrl+currentWeatherRequestUrl, nil)

	if err != nil {
		log.Println("Could not create request to external Weather API", err)
		return model.Weather{}, err
	}

	queryParams := request.URL.Query()

	queryParams.Add(apiKeyParam, apiKey)
	queryParams.Add(queryParam, cityName)
	queryParams.Add(airQualityParam, no)

	request.URL.RawQuery = queryParams.Encode()

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Println("Error while requesting external Weather API", err)
		return model.Weather{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error while closing Body Reader", err)
		}
	}(response.Body)

	var externalWeatherResponse model.ExternalWeatherResponse
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&externalWeatherResponse)

	if err != nil {
		log.Println("Error while decoding external Weather API response", err)
		return model.Weather{}, err
	}

	return model.Weather{
		Temperature: externalWeatherResponse.Current.TempC,
		Humidity:    externalWeatherResponse.Current.Humidity,
		Description: externalWeatherResponse.Current.Condition.Text,
	}, nil
}
