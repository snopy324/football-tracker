package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/joho/godotenv"
)

var (
	FootballDataAPIKey     string
	FootballDataBaseURL    string
	LineChannelAccessToken string
	LineAPIURL             string
	LineToID               string
)

func loadConfigFromEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}

	FootballDataAPIKey = os.Getenv("FOOTBALL_DATA_API_KEY")
	LineChannelAccessToken = os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	LineToID = os.Getenv("LINE_TO_ID")
}

func loadConfigFromAWS() {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error creating AWS session")
	}

	svc := secretsmanager.New(sess)

	secretName := os.Getenv("SECRET_ARN")
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error getting secret value")
	}

	secretString := *result.SecretString
	var config map[string]string
	json.Unmarshal([]byte(secretString), &config)

	FootballDataAPIKey = config["FOOTBALL_DATA_API_KEY"]
	LineChannelAccessToken = config["LINE_CHANNEL_ACCESS_TOKEN"]
	LineToID = config["LINE_TO_ID"]
}

func init() {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "test_function" {
		loadConfigFromAWS()
	} else {
		loadConfigFromEnv()
	}

	FootballDataBaseURL = "https://api.football-data.org/v4"
	LineAPIURL = "https://api.line.me/v2/bot/message/push"
}
