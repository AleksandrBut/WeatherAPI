package main

import (
	"WeatherAPI/client"
	"WeatherAPI/db"
	"WeatherAPI/model"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	db.CreateConnectionPool()
	defer db.CloseConnectionPool()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Mount("/api", api())

	log.Println("Application started")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error: ", err)
	}
}

func api() http.Handler {
	r := chi.NewRouter()

	r.Get("/weather", getWeatherByCity)
	r.Post("/subscribe", subscribeToWeatherUpdates)
	r.Get("/confirm/{token}", confirmEmailSubscription)
	r.Get("/unsubscribe/{token}", unsubscribeFromWeatherUpdates)
	r.Post("/test", test)

	return r
}

// TODO replace Write with Render
func test(w http.ResponseWriter, r *http.Request) {
	var subscription model.Subscription

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	err := decoder.Decode(&subscription)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Incorrect json structure"))

		return
	}

	subscription, err = db.CreateSubscription(subscription)

	if err != nil {
		//TODO replace with http.Error
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("DB error"))

		return
	}

	w.WriteHeader(http.StatusCreated)
	err = encoder.Encode(&subscription)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Could not encode response body"))

		return
	}
}

func getWeatherByCity(w http.ResponseWriter, r *http.Request) {
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

func subscribeToWeatherUpdates(w http.ResponseWriter, r *http.Request) {

}

func confirmEmailSubscription(w http.ResponseWriter, r *http.Request) {

}

func unsubscribeFromWeatherUpdates(w http.ResponseWriter, r *http.Request) {

}
