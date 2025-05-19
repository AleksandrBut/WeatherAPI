package main

import (
	"WeatherAPI/client"
	"WeatherAPI/db"
	"WeatherAPI/model"
	"WeatherAPI/token"
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

	return r
}

// TODO replace Write with Render
func test(w http.ResponseWriter, r *http.Request) {

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
	var subscription model.Subscription

	subscription.CityName = r.FormValue("city")
	subscription.Email = r.FormValue("email")
	subscription.Frequency = r.FormValue("frequency")

	if !model.IsSubscriptionValid(subscription) {
		log.Println("Invalid input for subscription: ", subscription)
		http.Error(w, "Invalid input", http.StatusBadRequest)

		return
	}

	if db.IsEmailAlreadySubscribed(subscription.Email) {
		log.Println("Email ", subscription.Email, " is already subscribed")
		http.Error(w, "Email already subscribed", http.StatusConflict)

		return
	}

	subscription.Token = token.GenerateToken()
	err := client.SendCreateSubscriptionEmail(&subscription)

	if err != nil {
		log.Println(err)
		http.Error(w, "Could not send confirmation email", http.StatusBadGateway)

		return
	}

	err = db.CreateSubscription(subscription)

	if err != nil {
		log.Println(err)
		http.Error(w, "DB error", http.StatusInternalServerError)
	}
}

func confirmEmailSubscription(w http.ResponseWriter, r *http.Request) {

}

func unsubscribeFromWeatherUpdates(w http.ResponseWriter, r *http.Request) {

}
