package application

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

type Handler func(app *Application, message *nsq.Message) error

var TestHandler Handler = func(app *Application, message *nsq.Message) error {
	fmt.Println(message)

	return nil
}
