package main

import (
	"log"
	"math/rand"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/mahikgot/gossip-glomers/internal/handler"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("generate", handler.Make(n, func(msg maelstrom.Message, requestBody, responseBody map[string]any) error {
		responseBody["id"] = rand.Int63()
		return nil
	}))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
