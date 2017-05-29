package test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/thoas/observr/application"
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

func GET(ctx context.Context, url string) *httptest.ResponseRecorder {
	return Request(ctx, "GET", url, nil)
}

func POST(ctx context.Context, url string, payload map[string]interface{}) *httptest.ResponseRecorder {
	return Request(ctx, "POST", url, payload)
}

func PATCH(ctx context.Context, url string, payload map[string]interface{}) *httptest.ResponseRecorder {
	return Request(ctx, "PATCH", url, payload)
}

func DELETE(ctx context.Context, url string) *httptest.ResponseRecorder {
	return Request(ctx, "DELETE", url, nil)
}

func Request(ctx context.Context, method string, url string, payload map[string]interface{}) *httptest.ResponseRecorder {
	var r *http.Request

	if payload != nil {
		body, err := json.Marshal(&payload)
		if err != nil {
			panic(err)
		}

		r, err = http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			panic(err)
		}

		r.Header.Set("Content-Type", "application/json")
	} else {
		r, _ = http.NewRequest(method, url, nil)
	}

	w := httptest.NewRecorder()

	router, err := web.Routes(ctx)

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, r)

	return w
}
