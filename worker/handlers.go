package worker

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
	"github.com/thoas/observr/broker"

	"context"
)

type Handler func(context.Context, []byte) error

// ErrorHandler is the error handler.
func ErrorHandler(handler Handler) broker.Handler {
	return func(ctx context.Context, message []byte) {
		err := handler(ctx, message)
		if err != nil {
			panic(err)
		}
	}
}

func UserCreatedHandler(ctx context.Context, message []byte) error {
	var result broker.UserCreatedEvent

	err := json.Unmarshal(message, &result)

	if err != nil {
		return err
	}

	spew.Dump(result)

	return nil
}
