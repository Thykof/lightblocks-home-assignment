package sqs

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQS struct {
	SQSServiceClient *sqs.SQS
	QueueUrl 	     string
}

func NewSQS(queueURL string) *SQS {
	region := os.Getenv("AWS_REGION")
	
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		log.Fatalf("Error creating session: %s", err)
	}

	svc := sqs.New(sess)

	return &SQS{
		SQSServiceClient: svc,
		QueueUrl: queueURL,
	}
}

func (s *SQS) SendMessage(message string, ) error {
	_, err := s.SQSServiceClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    aws.String(s.QueueUrl),
	})
	
	return err
}

func (s *SQS) ReceiveMessage() (*sqs.Message, error) {
	result, err := s.SQSServiceClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(s.QueueUrl),
		MaxNumberOfMessages: aws.Int64(1),
	})
	if err != nil {
		return nil, err
	}

	if len(result.Messages) == 0 {
		return nil, nil
	}

	return result.Messages[0], nil
}

func (s *SQS) PollMessages(chn chan<- *sqs.Message) {

	for {
		message, err := s.ReceiveMessage()
		if err != nil {
			log.Printf("failed to fetch sqs message %v", err)
		}

		chn <- message
	}
}