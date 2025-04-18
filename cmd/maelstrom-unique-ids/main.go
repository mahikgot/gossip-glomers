package main

import (
	"log"
	"math/rand"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/mahikgot/gossip-glomers/internal/handler"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("generate", handler.Make(n, func(msgBody map[string]any) (map[string]any, error) {
		msgBody["id"] = rand.Int63()
		return msgBody, nil
	}))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
