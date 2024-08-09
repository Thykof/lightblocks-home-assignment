package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thykof/lightblocks-home-assignment/pkg/input"
	"github.com/thykof/lightblocks-home-assignment/pkg/sqs"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	inputFilePath := os.Getenv("INPUT_FILE_PATH")
	queueURL := os.Getenv("AWS_QUEUE_URL")

	s := sqs.NewSQS(queueURL)

	messages := input.GetInputMessages(inputFilePath)

	for _, message := range messages {
		log.Printf("Sending message: %s\n", message)
		if err := s.SendMessage(message); err != nil {
			log.Fatalf("Error sending message: %s", err)
		}
	}
}