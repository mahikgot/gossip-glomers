package handler

import (
	"encoding/json"
	"errors"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type LogicFunc func(requestBody, responseBody *map[string]any) error

func Make(n *maelstrom.Node, fn LogicFunc) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		requestBody := make(map[string]any)
		if err := json.Unmarshal(msg.Body, &requestBody); err != nil {
			return err
		}

		msgType, ok := requestBody["type"].(string)
		if !ok {
			return errors.New("handler.Make err: msg type is not string")
		}
		responseBody := make(map[string]any)
		responseBody["type"] = msgType + "_ok"

		err := fn(&requestBody, &responseBody)
		if err != nil {
			return err
		}

		return n.Reply(msg, responseBody)
	}
}
