package server

import (
	"errors"
	"log"
	"strings"
)

func (s Server) Handle(message string) {
	log.Printf("Received message: %s", message)

	messageSplit := strings.Split(message, "(")
	if len(messageSplit) != 2 {
		log.Printf("Invalid message: %s", message)
		return
	}

	functionName := messageSplit[0]
	params := messageSplit[1]

	switch functionName {
		case "addItem":
			s.addItem(params)
		case "deleteItem":
			s.deleteItem(params)
		case "getItem":
			s.GetItem(params)
		case "getAllItems":
			s.getAllItems()
		default:
			log.Printf("Invalid function: %s", functionName)
	}
}

func (s *Server) addItem(params string) {
	key, err := s.parseKey(params)
	if err != nil {
		log.Printf("Error parsing key: %s", err)
		return
	}

	value, err := s.parseValue(params)
	if err != nil {
		log.Printf("Error parsing value: %s", err)
		return
	}

	err = s.Data.Set(key, value)
	if err != nil {
		log.Printf("Error adding item: %s", err)
		return
	}

	log.Printf("Adding item: %s => %s", key, value)
}

func (s *Server) deleteItem(params string) {
	key, err := s.parseKey(params)
	if err != nil {
		log.Printf("Error parsing key: %s", err)
		return
	}

	s.Data.Delete(key)
	log.Printf("Deleting item: %s", key)
}

func (s *Server) GetItem(params string) {
	key, err := s.parseKey(params)
	if err != nil {
		log.Printf("Error parsing key: %s", err)
		return
	}

	value := s.Data.Get(key)
	log.Printf("Item: %s => %s", key, value)
}

func (s *Server) getAllItems() {
	log.Printf("Getting all items")
	for _, pair := range s.Data.GetAll() {
		log.Printf("Item: %s => %s", pair.Key, pair.Value)
	}
}

func (s *Server) parseKey(params string) (string, error) {
	paramsSplit := strings.Split(params, "'")
	if len(paramsSplit) < 1 {
		return "", errors.New("invalid params")
	}

	return paramsSplit[1], nil
}

func (s *Server) parseValue(params string) (string, error) {
	paramsSplit := strings.Split(params, "'")
	if len(paramsSplit) < 3 {
		return "", errors.New("invalid params")
	}

	return paramsSplit[3], nil
}