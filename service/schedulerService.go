package service

import (
	"WeatherAPI/client"
	"WeatherAPI/db"
	"WeatherAPI/model"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

const kyivLocation = "Europe/Kyiv"

func InitScheduler() {
	location, _ := time.LoadLocation(kyivLocation)
	scheduler := cron.New(cron.WithLocation(location))

	_, _ = scheduler.AddFunc("* * * * *", hourlyUpdate)
	_, _ = scheduler.AddFunc("*/10 * * * *", dailyUpdate)

	scheduler.Start()
}

func hourlyUpdate() {
	handleWeatherUpdate(model.Hourly)
}

func dailyUpdate() {
	handleWeatherUpdate(model.Daily)
}

func handleWeatherUpdate(frequency string) {
	subscriptions, err := db.GetActiveSubscriptions(frequency)

	if err != nil {
		log.Println("DB error in hourly scheduler: ", err)
	}

	for _, s := range subscriptions {
		weather, err := client.GetWeatherByCity(s.CityName)

		if err != nil {
			log.Println("Could not get external weather data: ", err)
			continue
		}

		go func() {
			err := client.SendWeatherUpdateEmail(&s, &weather)

			if err != nil {
				log.Println("Could not send weather update email: ", err)
			}
		}()
	}
}
