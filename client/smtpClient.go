package client

import (
	"WeatherAPI/model"
	_ "github.com/joho/godotenv/autoload"
	"net/smtp"
	"os"
	"strconv"
	"strings"
)

const smtpHost = "smtpHost"
const smtpPort = "smtpPort"
const from = "from"
const password = "password"
const to = "to"
const message = "message"

const baseUrl = "http://localhost:8080/api/"
const confirmSubscriptionBaseUrl = "confirm/"
const unsubscribeBaseUrl = "unsubscribe/"

const emailSubscriptionSubjectPart = "Subject: Weather update confirmation\n\n"
const emailSubscriptionHiPart = "Hi!\n\nYou are going to subscribe to weather updates in "
const emailSubscriptionOnPart = " on "
const emailSubscriptionConfirmPart = " basis.\nPlease confirm your subscription:\n"

const emailWeatherUpdateSubjectPart = "Subject: Weather update\n\n"
const emailWeatherUpdateHiPart = "Hi!\n\nYour weather update:\n"
const emailWeatherUpdateCityPart = "\nCity: "
const emailWeatherUpdateTemperaturePart = "\nTemperature: "
const emailWeatherUpdateHumidityPart = "\nHumidity: "
const emailWeatherUpdateDescriptionPart = "\nDescription: "
const emailWeatherUpdateFrequencyPart = "\nWeather update frequency: "
const emailWeatherUpdateUnsubscribePart = "\n\nUnsubscribe from weather updates:\n"

func SendCreateSubscriptionEmail(subscription *model.Subscription) error {
	confirmSubscriptionUrl := baseUrl + confirmSubscriptionBaseUrl + subscription.Token
	var stringBuilder strings.Builder

	stringBuilder.WriteString(emailSubscriptionSubjectPart)
	stringBuilder.WriteString(emailSubscriptionHiPart)
	stringBuilder.WriteString(subscription.CityName)
	stringBuilder.WriteString(emailSubscriptionOnPart)
	stringBuilder.WriteString(subscription.Frequency)
	stringBuilder.WriteString(emailSubscriptionConfirmPart)
	stringBuilder.WriteString(confirmSubscriptionUrl)

	config := getConfigMap(subscription, stringBuilder.String())

	return sendEmail(config)
}

func SendWeatherUpdateEmail(subscription *model.Subscription, weather *model.Weather) error {
	unsubscribeUrl := baseUrl + unsubscribeBaseUrl + subscription.Token
	var stringBuilder strings.Builder

	stringBuilder.WriteString(emailWeatherUpdateSubjectPart)
	stringBuilder.WriteString(emailWeatherUpdateHiPart)
	stringBuilder.WriteString(emailWeatherUpdateCityPart)
	stringBuilder.WriteString(subscription.CityName)
	stringBuilder.WriteString(emailWeatherUpdateTemperaturePart)
	stringBuilder.WriteString(strconv.FormatFloat(weather.Temperature, 'f', 2, 64))
	stringBuilder.WriteString(emailWeatherUpdateHumidityPart)
	stringBuilder.WriteString(strconv.Itoa(weather.Humidity))
	stringBuilder.WriteString(emailWeatherUpdateDescriptionPart)
	stringBuilder.WriteString(weather.Description)
	stringBuilder.WriteString(emailWeatherUpdateFrequencyPart)
	stringBuilder.WriteString(subscription.Frequency)
	stringBuilder.WriteString(emailWeatherUpdateUnsubscribePart)
	stringBuilder.WriteString(unsubscribeUrl)

	config := getConfigMap(subscription, stringBuilder.String())

	return sendEmail(config)
}

func getConfigMap(subscription *model.Subscription, msg string) map[string]string {
	config := make(map[string]string)

	config[smtpHost] = os.Getenv("SMTP_HOST")
	config[smtpPort] = os.Getenv("SMTP_PORT")
	config[from] = os.Getenv("SMTP_FROM")
	config[password] = os.Getenv("SMTP_PASSWORD")
	config[to] = subscription.Email
	config[message] = msg

	return config
}

func sendEmail(config map[string]string) error {
	auth := smtp.PlainAuth("", config[from], config[password], config[smtpHost])

	return smtp.SendMail(
		config[smtpHost]+":"+config[smtpPort],
		auth,
		config[from],
		[]string{config[to]},
		[]byte(config[message]),
	)
}
