Weather API

Written in Go, database is Postgres, startup migration is done using Goose, smtp server is gmail, scheduling is implemented using cron, routing - using chi

Needs .env file in root to run. For example:


GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://postgres:postgres@weather_api_postgres_db:5432/weather_api
GOOSE_MIGRATION_DIR=./db/migration


EXTERNAL_WEATHER_API_KEY=key


SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_PASSWORD=password
SMTP_FROM=email


Operations:


GET /api/weather?city=Odesa - returns current weather for specified city. Details: uses external api.weatherapi.com

POST /api/subscribe - creates a subscription to specified city, with specified frequency. Sends confirmation email to specified email

GET /api/confirm/{token} - confirms subscription. Details: after this request a weather update will be sent to previously specified email with previously specified frequency

GET /api/unsubscribe/{token} - unsubscribes from weather updates, removes all subscription data from db


Confirmation email example:

![image](https://github.com/user-attachments/assets/0b50db15-326e-4e1f-885e-ec5f5234a7a0)

Weather update email example:

![image](https://github.com/user-attachments/assets/47f0af74-1981-455c-87bf-01ab69b75c59)

