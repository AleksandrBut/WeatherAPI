package service

import (
	"WeatherAPI/client"
	"encoding/json"
	"net/http"
)

func GetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	cityName := r.URL.Query().Get("city")

	if cityName == "" {
		http.Error(w, "Request parameter 'city' must be specified", http.StatusBadRequest)
		return
	}

	weather, err := client.GetWeatherByCity(cityName)

	if err != nil {
		http.Error(w, "Error while requesting external Weather API: "+err.Error(), http.StatusBadGateway)
		return
	}

	if err = json.NewEncoder(w).Encode(&weather); err != nil {
		http.Error(w, "Could not encode response body", http.StatusInternalServerError)
		return
	}
}
