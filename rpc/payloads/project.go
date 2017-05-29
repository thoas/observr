package payloads

import (
	"net/http"

	"github.com/mholt/binding"
)

type ProjectCreate struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (p *ProjectCreate) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.Name: binding.Field{
			Form:     "name",
			Required: true,
		},
		&p.URL: binding.Field{
			Form:     "url",
			Required: true,
		},
	}
}
