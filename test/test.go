package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/thoas/observr/application"
	"github.com/thoas/observr/store/models"
	"github.com/thoas/observr/web"
)

func Setup(fn func(ctx context.Context)) {
	ctx, err := application.Load(os.Getenv("OBSERVR_CONF"))

	if err != nil {
		panic(err)
	}

	fn(ctx)

	err = application.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}

type Request struct {
	Method  string
	URL     string
	Data    map[string]interface{}
	Headers map[string]interface{}
	User    *models.User
}

func GET(ctx context.Context, req *Request) *httptest.ResponseRecorder {
	req.Method = "GET"

	return request(ctx, req)
}

func POST(ctx context.Context, req *Request) *httptest.ResponseRecorder {
	req.Method = "POST"

	return request(ctx, req)
}

func PATCH(ctx context.Context, req *Request) *httptest.ResponseRecorder {
	req.Method = "PATCH"

	return request(ctx, req)
}

func DELETE(ctx context.Context, req *Request) *httptest.ResponseRecorder {
	req.Method = "DELETE"

	return request(ctx, req)
}

func request(ctx context.Context, req *Request) *httptest.ResponseRecorder {
	var r *http.Request

	if req.Data != nil {
		body, err := json.Marshal(&req.Data)
		if err != nil {
			panic(err)
		}

		r, err = http.NewRequest(req.Method, req.URL, bytes.NewBuffer(body))
		if err != nil {
			panic(err)
		}

		r.Header.Set("Content-Type", "application/json")
	} else {
		r, _ = http.NewRequest(req.Method, req.URL, nil)
	}
	if req.User != nil {
		r.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", req.User.ApiKey))
	}

	w := httptest.NewRecorder()

	router, err := web.Routes(ctx)

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, r)

	return w
}
