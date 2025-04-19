package handler

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func readRequest(msg *maelstrom.Message) (map[string]any, error) {
	body := make(map[string]any)
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return nil, err
	}

	return body, nil
}

type LogicFunc func(requestBody, responseBody map[string]any) error

func Make(n *maelstrom.Node, fn LogicFunc) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		requestBody, err := readRequest(&msg)
		if err != nil {
			return err
		}

		responseBody := make(map[string]any)
		responseBody["type"] = msg.Type() + "_ok"

		err = fn(requestBody, responseBody)
		if err != nil {
			return err
		}

		return n.Reply(msg, responseBody)
	}
}
