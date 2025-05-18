package model

type Subscription struct {
	Id        int    `json:"id,omitempty"`
	CityName  string `json:"cityName,omitempty"`
	Email     string `json:"email,omitempty"`
	Frequency string `json:"frequency,omitempty"`
}

type Weather struct {
	Temperature float64 `json:"temperature,omitempty"`
	Humidity    int     `json:"humidity,omitempty"`
	Description string  `json:"description,omitempty"`
}

type ExternalWeatherResponse struct {
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		Humidity int `json:"humidity"`
	} `json:"current"`
}
