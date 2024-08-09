package handler

import "log"

func Handle(message string) {
	log.Printf("Received message: %s", message)
}