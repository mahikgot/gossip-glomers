package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/mahikgot/gossip-glomers/internal/handler"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("echo", handler.Make(n, func(msgBody map[string]any) (map[string]any, error) {
		return msgBody, nil
	}))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
