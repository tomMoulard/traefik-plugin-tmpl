package traefik_plugin_tmpl_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	tmpl "github.com/tomMoulard/traefik-plugin-tmpl"
)

func TestTmpl(t *testing.T) {
	t.Parallel()

	cfg := tmpl.CreateConfig()

	var nextCount int

	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		nextCount++
	})

	handler, err := tmpl.New(context.Background(), next, cfg, "tmpl")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://localhost/tmpl.go", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("invalid recorder status code, expected: %d, got: %d", http.StatusOK, recorder.Code)
	}

	if nextCount != 1 {
		t.Errorf("invalid next count, expected: %d, got: %d", 1, nextCount)
	}
}
