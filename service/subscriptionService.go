package service

import (
	"WeatherAPI/client"
	"WeatherAPI/db"
	"WeatherAPI/model"
	"WeatherAPI/token"
	"log"
	"net/http"
)

func SubscribeToWeatherUpdates(w http.ResponseWriter, r *http.Request) {
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
