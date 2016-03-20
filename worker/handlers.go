package worker

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/context"

	"github.com/nsqio/go-nsq"
)

type Handler func(message *nsq.Message, ctx context.Context) error

var TestHandler Handler = func(message *nsq.Message, ctx context.Context) error {
	var result struct {
		Foo string `json:"foo"`
	}

	err := json.Unmarshal(message.Body, &result)

	if err != nil {
		return err
	}

	fmt.Println(result.Foo)

	return nil
}
