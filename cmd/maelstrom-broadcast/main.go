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

	n.Handle("broadcast", handler.Make(n, func(msgBody map[string]any) (map[string]any, error) {
		message, ok := msgBody["message"].(float64)
		if !ok {
			return nil, errors.New("broadcastHandle: body message type not int")
		}
		sl = append(sl, message)

		responseBody := make(map[string]any)
		return responseBody, nil
	}))

	n.Handle("read", handler.Make(n, func(msgBody map[string]any) (map[string]any, error) {
		responseBody := make(map[string]any)
		responseBody["messages"] = sl
		return responseBody, nil
	}))

	n.Handle("topology", handler.Make(n, func(msgBody map[string]any) (map[string]any, error) {
		responseBody := make(map[string]any)
		return responseBody, nil
	}))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
