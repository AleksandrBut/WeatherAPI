package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Mount("/api", Api())

	_ = http.ListenAndServe(":8080", r)
}

func Api() http.Handler {
	r := chi.NewRouter()

	r.Get("/weather", GetWeatherByCity)
	r.Post("/subscribe", SubscribeToWeatherUpdates)
	r.Get("/confirm/{token}", ConfirmEmailSubscription)
	r.Get("/unsubscribe/{token}", UnsubscribeFromWeatherUpdates)

	return r
}

func GetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Test weather response"))
}

func SubscribeToWeatherUpdates(w http.ResponseWriter, r *http.Request) {

}

func ConfirmEmailSubscription(w http.ResponseWriter, r *http.Request) {

}

func UnsubscribeFromWeatherUpdates(w http.ResponseWriter, r *http.Request) {

}
