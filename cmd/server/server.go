package main

import (
	"log"
	"os"

	awssqs "github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joho/godotenv"
	serverpkg "github.com/thykof/lightblocks-home-assignment/int/server"
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

	server := serverpkg.NewServer()

	chnMessages := make(chan *awssqs.Message)
	go s.PollMessages(chnMessages)

	for message := range chnMessages {
		server.Handle(*message.Body)
		s.DeleteMessage(message.ReceiptHandle)
	}
}
