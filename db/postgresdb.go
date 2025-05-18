package db

import (
	"WeatherAPI/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var pool *pgxpool.Pool

func CreateConnectionPool() {
	pool, _ = pgxpool.New(context.Background(), os.Getenv("GOOSE_DBSTRING"))
}

func CreateSubscription(subscription model.Subscription) (model.Subscription, error) {
	err := pool.QueryRow(
		context.Background(),
		"insert into subscription(cityname, email, frequency) values ($1, $2, $3) returning id",
		subscription.CityName,
		subscription.Email,
		subscription.Frequency,
	).Scan(&subscription.Id)

	if err != nil {
		return model.Subscription{}, err
	}

	return subscription, nil
}

func CloseConnectionPool() {
	pool.Close()
}
