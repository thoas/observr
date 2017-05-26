package payloads

import (
	"net/http"

	"github.com/mholt/binding"
)

type UserCreate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *UserCreate) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.Username: binding.Field{
			Form:     "username",
			Required: true,
		},
		&p.Email: binding.Field{
			Form:     "email",
			Required: true,
		},
		&p.Password: binding.Field{
			Form:     "password",
			Required: true,
		},
	}
}
