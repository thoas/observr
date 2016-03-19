package application

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
)

type Handler func(app *Application, message *nsq.Message) error

var TestHandler Handler = func(app *Application, message *nsq.Message) error {
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
