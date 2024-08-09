package main

import (
	"log"
	"os"

	awssqs "github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joho/godotenv"
	"github.com/thykof/lightblocks-home-assignment/int/handler"
	"github.com/thykof/lightblocks-home-assignment/pkg/sqs"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	queueURL := os.Getenv("AWS_QUEUE_URL")

	s := sqs.NewSQS(queueURL)

	chnMessages := make(chan *awssqs.Message)
	go s.PollMessages(chnMessages)

	for message := range chnMessages {
		if message != nil {
			handler.Handle(*message.Body)
		}
	}
}
