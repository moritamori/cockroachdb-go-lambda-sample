package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jackc/pgx/v4"
)

func HandleRequest(ctx context.Context) (string, error) {
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=verify-full", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("HOSTNAME"), os.Getenv("PORT"), os.Getenv("DB_NAME"))
	config, err := pgx.ParseConfig(url)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := pgx.ConnectConfig(context.Background(), config)
	defer conn.Close(context.Background())
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	var now time.Time
	err = conn.QueryRow(ctx, "SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	return fmt.Sprintf("%s", now), nil
}

func main() {
	lambda.Start(HandleRequest)
}
