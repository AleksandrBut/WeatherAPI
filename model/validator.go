package model

const Hourly = "hourly"
const Daily = "daily"

func IsSubscriptionValid(subscription *Subscription) bool {
	return subscription.CityName != "" &&
		subscription.Email != "" &&
		subscription.Frequency != "" &&
		(subscription.Frequency == Hourly || subscription.Frequency == Daily)
}

func IsConfirmationTokenValid(token *string) bool {
	return *token != ""
}
