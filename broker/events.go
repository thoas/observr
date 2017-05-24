package broker

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type UserCreatedEvent struct {
	Username string `json:"username"`
}

func (UserCreatedEvent) Name() string {
	return "observr.user.created"
}

func (e *UserCreatedEvent) ToBytes() ([]byte, error) {
	msg, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot marshall %T as json payload", e)
	}

	return msg, nil
}
