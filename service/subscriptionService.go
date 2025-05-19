package service

import (
	"WeatherAPI/client"
	"WeatherAPI/db"
	"WeatherAPI/model"
	"WeatherAPI/token"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

const tokenParam = "token"

func SubscribeToWeatherUpdates(w http.ResponseWriter, r *http.Request) {
	var subscription model.Subscription

	subscription.CityName = r.FormValue("city")
	subscription.Email = r.FormValue("email")
	subscription.Frequency = r.FormValue("frequency")

	if !model.IsSubscriptionValid(&subscription) {
		log.Println("Invalid input for subscription: ", subscription)
		http.Error(w, "Invalid input", http.StatusBadRequest)

		return
	}

	if db.IsEmailAlreadySubscribed(&subscription.Email) {
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

	err = db.CreateSubscription(&subscription)

	if err != nil {
		log.Println(err)
		http.Error(w, "DB error", http.StatusInternalServerError)
	}
}

func ConfirmEmailSubscription(w http.ResponseWriter, r *http.Request) {
	confirmationToken := getTokenFromPath(r)

	if !model.IsConfirmationTokenValid(&confirmationToken) {
		log.Println("Invalid token ", confirmationToken)
		http.Error(w, "Invalid token", http.StatusBadRequest)
	}

	subscriptionId, err := db.GetSubscriptionIdByToken(&confirmationToken)

	if subscriptionId == 0 {
		log.Println("Token not found ", confirmationToken)
		http.Error(w, "Token not found", http.StatusNotFound)

		return
	}

	if err != nil {
		log.Println("DB error")
		http.Error(w, "DB error", http.StatusInternalServerError)

		return
	}

	if err = db.SetSubscriptionActiveById(&subscriptionId); err != nil {
		log.Println("DB error")
		http.Error(w, "DB error", http.StatusInternalServerError)
	}
}

func Unsubscribe(w http.ResponseWriter, r *http.Request) {
	confirmationToken := getTokenFromPath(r)

	if !model.IsConfirmationTokenValid(&confirmationToken) {
		log.Println("Invalid token ", confirmationToken)
		http.Error(w, "Invalid token", http.StatusBadRequest)
	}

	subscriptionId, err := db.GetSubscriptionIdByToken(&confirmationToken)

	if subscriptionId == 0 {
		log.Println("Token not found ", confirmationToken)
		http.Error(w, "Token not found", http.StatusNotFound)

		return
	}

	if err != nil {
		log.Println("DB error")
		http.Error(w, "DB error", http.StatusInternalServerError)

		return
	}

	if err = db.DeleteSubscriptionById(&subscriptionId); err != nil {
		log.Println("DB error")
		http.Error(w, "DB error", http.StatusInternalServerError)
	}
}

func getTokenFromPath(r *http.Request) string {
	return chi.URLParam(r, tokenParam)
}
