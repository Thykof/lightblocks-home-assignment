package server

import "github.com/thykof/lightblocks-home-assignment/pkg/orderedmap"

type Server struct {
	Data *orderedmap.OrderedMap
}

func NewServer() *Server {
	return &Server{
		Data: orderedmap.NewOrderedMap(),
	}
}