package model

const hourly = "hourly"
const daily = "daily"

func IsSubscriptionValid(subscription *Subscription) bool {
	return subscription.CityName != "" &&
		subscription.Email != "" &&
		subscription.Frequency != "" &&
		(subscription.Frequency == hourly || subscription.Frequency == daily)
}

func IsConfirmationTokenValid(token *string) bool {
	return *token != ""
}
