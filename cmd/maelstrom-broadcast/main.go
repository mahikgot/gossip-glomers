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
	messages map[float64]string
	mu       sync.Mutex
	topology []string
}

func main() {
	server := &Server{
		node:     maelstrom.NewNode(),
		messages: make(map[float64]string),
	}

	server.node.Handle("broadcast", server.Broadcast())
	server.node.Handle("read", server.Read())
	server.node.Handle("topology", server.Topology())
	server.node.Handle("broadcast_ok", server.Noop())

	server.Run()
}

func (s *Server) Broadcast() maelstrom.HandlerFunc {
	return handler.Make(s.node, func(msg maelstrom.Message, requestBody, responseBody map[string]any) error {
		message, ok := requestBody["message"].(float64)
		if !ok {
			return errors.New("broadcastHandle: body message type not int")
		}

		s.mu.Lock()
		if _, ok := s.messages[message]; ok {
			s.mu.Unlock()
			return nil
		}
		s.messages[message] = ""
		s.mu.Unlock()

		broadcastBody := make(map[string]any)
		broadcastBody["type"] = "broadcast"
		broadcastBody["message"] = message

		for _, id := range s.topology {
			if id == s.node.ID() || id == msg.Src {
				continue
			}
			s.node.Send(id, broadcastBody)
		}

		return nil
	})
}

func (s *Server) Read() maelstrom.HandlerFunc {
	return handler.Make(s.node, func(msg maelstrom.Message, requestBody, responseBody map[string]any) error {
		s.mu.Lock()
		defer s.mu.Unlock()
		responseBody["messages"] = handler.MapKeysToSlice(s.messages)
		return nil
	})
}

func (s *Server) Topology() maelstrom.HandlerFunc {
	return handler.Make(s.node, func(msg maelstrom.Message, requestBody, responseBody map[string]any) error {
		s.topology = s.node.NodeIDs()
		return nil
	})
}

func (s *Server) Noop() maelstrom.HandlerFunc {
	return handler.Nothing()
}

func (s *Server) Run() {
	if err := s.node.Run(); err != nil {
		log.Fatal(err)
	}
}
