package main

import (
	"WeatherAPI/db"
	"WeatherAPI/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	db.CreateConnectionPool()
	defer db.CloseConnectionPool()

	service.InitScheduler()

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

	return r
}

func getWeatherByCity(w http.ResponseWriter, r *http.Request) {
	service.GetWeatherByCity(w, r)
}

func subscribeToWeatherUpdates(w http.ResponseWriter, r *http.Request) {
	service.SubscribeToWeatherUpdates(w, r)
}

func confirmEmailSubscription(w http.ResponseWriter, r *http.Request) {
	service.ConfirmEmailSubscription(w, r)
}

func unsubscribeFromWeatherUpdates(w http.ResponseWriter, r *http.Request) {
	service.Unsubscribe(w, r)
}
