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

type LogicFunc func(msg maelstrom.Message, requestBody, responseBody map[string]any) error

func Make(n *maelstrom.Node, fn LogicFunc) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		requestBody, err := readRequest(&msg)
		if err != nil {
			return err
		}

		responseBody := make(map[string]any)
		responseBody["type"] = msg.Type() + "_ok"

		err = fn(msg, requestBody, responseBody)
		if err != nil {
			return err
		}

		return n.Reply(msg, responseBody)
	}
}

func Nothing() maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		return nil
	}
}

func MapKeysToSlice[K comparable, V any](mapp map[K]V) []K {
	keys := make([]K, 0, len(mapp))
	for k := range mapp {
		keys = append(keys, k)
	}
	return keys
}
