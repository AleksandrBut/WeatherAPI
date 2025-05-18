package model

type Subscription struct {
	Id        int    `json:"id,omitempty"`
	CityName  string `json:"cityName,omitempty"`
	Email     string `json:"email,omitempty"`
	Frequency string `json:"frequency,omitempty"`
}
