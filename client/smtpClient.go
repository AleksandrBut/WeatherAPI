package client

import (
	"WeatherAPI/model"
	_ "github.com/joho/godotenv/autoload"
	"net/smtp"
	"os"
	"strings"
)

const confirmSubscriptionBaseUrl = "http://localhost:8080/api/confirm/"
const emailSubjectPart = "Subject: Weather update confirmation\n\n"
const emailHiPart = "Hi!\n\nYou are going to subscribe to weather updates in "
const emailOnPart = " on "
const emailConfirmPart = " basis.\nPlease confirm your subscription:\n"

func SendCreateSubscriptionEmail(subscription *model.Subscription) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	from := os.Getenv("SMTP_FROM")
	password := os.Getenv("SMTP_PASSWORD")
	to := []string{subscription.Email}
	confirmSubscriptionUrl := confirmSubscriptionBaseUrl + subscription.Token

	var stringBuilder strings.Builder

	stringBuilder.WriteString(emailSubjectPart)
	stringBuilder.WriteString(emailHiPart)
	stringBuilder.WriteString(subscription.CityName)
	stringBuilder.WriteString(emailOnPart)
	stringBuilder.WriteString(subscription.Frequency)
	stringBuilder.WriteString(emailConfirmPart)
	stringBuilder.WriteString(confirmSubscriptionUrl)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	return smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
		to,
		[]byte(stringBuilder.String()),
	)
}
