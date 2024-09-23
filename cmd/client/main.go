package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/codescalersinternships/Datetime-Client-Marwan-Radwan/pkg/client"
	"github.com/joho/godotenv"
)

func main() {
	var isJson bool
	flag.BoolVar(&isJson, "j", false, "use the json endpoint")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseUrl := os.Getenv("BASE_URL")

	httpClient := client.NewClient(baseUrl)

	operation := func() error {
		dateTime, err := httpClient.GetDateTime(isJson)
		if err != nil {
			return err
		}
		fmt.Println("Current DateTime:", dateTime)
		return nil
	}

	if err := client.Retry(operation); err != nil {
		log.Fatalf("Failed to get datetime: %v", err)
	}
}
