package main

import (
	"errors"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/mahikgot/gossip-glomers/internal/handler"
)

func main() {
	n := maelstrom.NewNode()
	sl := make([]float64, 0, 1024)

	n.Handle("broadcast", handler.Make(n, func(requestBody, responseBody *map[string]any) error {
		message, ok := (*requestBody)["message"].(float64)
		if !ok {
			return errors.New("broadcastHandle: body message type not int")
		}
		sl = append(sl, message)

		return nil
	}))

	n.Handle("read", handler.Make(n, func(requestBody, responseBody *map[string]any) error {
		(*responseBody)["messages"] = sl
		return nil
	}))

	n.Handle("topology", handler.Make(n, func(requestBody, responseBody *map[string]any) error {
		return nil
	}))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
