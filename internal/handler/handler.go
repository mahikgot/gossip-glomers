package handler

import (
	"encoding/json"
	"errors"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type LogicFunc func(msgBody map[string]any) (map[string]any, error)

func Make(n *maelstrom.Node, fn LogicFunc) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		msgBody := make(map[string]any)
		if err := json.Unmarshal(msg.Body, &msgBody); err != nil {
			return err
		}

		responseBody, err := fn(msgBody)
		if err != nil {
			return err
		}

		msgType, ok := msgBody["type"].(string)
		if !ok {
			return errors.New("handler.Make err: msg type is not string")
		}
		responseBody["type"] = msgType + "_ok"

		return n.Reply(msg, responseBody)
	}
}
