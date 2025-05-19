package db

import (
	"WeatherAPI/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var pool *pgxpool.Pool

func CreateConnectionPool() {
	pool, _ = pgxpool.New(context.Background(), os.Getenv("GOOSE_DBSTRING"))

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("Could not connect to database: ", err)
	}
}

func CreateSubscription(subscription model.Subscription) error {
	_, err := pool.Exec(
		context.Background(),
		"insert into subscription(cityname, email, frequency, token) values ($1, $2, $3, $4)",
		subscription.CityName,
		subscription.Email,
		subscription.Frequency,
		subscription.Token,
	)

	return err
}

func IsEmailAlreadySubscribed(email string) bool {
	var exists bool
	_ = pool.QueryRow(
		context.Background(),
		"select exists (select 1 from subscription where email = $1) ",
		email,
	).Scan(&exists)

	return exists
}

func CloseConnectionPool() {
	pool.Close()
}
