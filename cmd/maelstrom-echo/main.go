package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/mahikgot/gossip-glomers/internal/handler"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("echo", handler.Make(n, func(requestBody, responseBody map[string]any) error {
		responseBody["echo"] = requestBody["echo"]
		return nil
	}))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
