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

func CreateSubscription(subscription *model.Subscription) error {
	_, err := pool.Exec(
		context.Background(),
		"insert into subscription(city_name, email, frequency, token) values ($1, $2, $3, $4)",
		subscription.CityName,
		subscription.Email,
		subscription.Frequency,
		subscription.Token,
	)

	return err
}

func IsEmailAlreadySubscribed(email *string) bool {
	var exists bool
	_ = pool.QueryRow(
		context.Background(),
		"select exists (select 1 from subscription where email = $1) ",
		email,
	).Scan(&exists)

	return exists
}

func GetSubscriptionIdByToken(token *string) (int, error) {
	var subscriptionId int

	err := pool.QueryRow(
		context.Background(),
		"select id from subscription where token = $1",
		token,
	).Scan(&subscriptionId)

	return subscriptionId, err
}

func SetSubscriptionActiveById(id *int) error {
	_, err := pool.Exec(
		context.Background(),
		"update subscription set is_active = true where id = $1",
		id,
	)

	return err
}

func CloseConnectionPool() {
	pool.Close()
}

func DeleteSubscriptionById(id *int) error {
	_, err := pool.Exec(
		context.Background(),
		"delete from subscription where id = $1",
		id,
	)

	return err
}

func GetActiveSubscriptions(frequency string) ([]model.Subscription, error) {
	rows, err := pool.Query(
		context.Background(),
		"select id, city_name, email, token, frequency from subscription where is_active = true and frequency = $1",
		frequency,
	)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var subscriptions []model.Subscription

	for rows.Next() {
		var subscription model.Subscription

		if err = rows.Scan(
			&subscription.Id,
			&subscription.CityName,
			&subscription.Email,
			&subscription.Token,
			&subscription.Frequency,
		); err != nil {
			return nil, err
		}

		subscriptions = append(subscriptions, subscription)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return subscriptions, nil
}
