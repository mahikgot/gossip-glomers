package main

import (
	"errors"
	"log"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/mahikgot/gossip-glomers/internal/handler"
)

type Server struct {
	node     *maelstrom.Node
	messages []float64
	mu       sync.Mutex
}

func main() {
	server := &Server{
		node:     maelstrom.NewNode(),
		messages: make([]float64, 0, 1024),
	}

	server.node.Handle("broadcast", server.Broadcast())
	server.node.Handle("read", server.Read())
	server.node.Handle("topology", server.Topology())

	server.Run()
}

func (s *Server) Broadcast() maelstrom.HandlerFunc {
	return handler.Make(s.node, func(requestBody, responseBody map[string]any) error {
		message, ok := requestBody["message"].(float64)
		if !ok {
			return errors.New("broadcastHandle: body message type not int")
		}

		defer s.mu.Unlock()
		s.mu.Lock()
		s.messages = append(s.messages, message)

		return nil
	})
}

func (s *Server) Read() maelstrom.HandlerFunc {
	return handler.Make(s.node, func(requestBody, responseBody map[string]any) error {
		responseBody["messages"] = s.messages
		return nil
	})
}

func (s *Server) Topology() maelstrom.HandlerFunc {
	return handler.Make(s.node, func(requestBody, responseBody map[string]any) error {
		return nil
	})
}

func (s *Server) Run() {
	if err := s.node.Run(); err != nil {
		log.Fatal(err)
	}
}
