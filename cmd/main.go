package main

import (
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
	_, _ = w.Write([]byte("Test weather response"))
}

func subscribeToWeatherUpdates(w http.ResponseWriter, r *http.Request) {

}

func confirmEmailSubscription(w http.ResponseWriter, r *http.Request) {

}

func unsubscribeFromWeatherUpdates(w http.ResponseWriter, r *http.Request) {

}
