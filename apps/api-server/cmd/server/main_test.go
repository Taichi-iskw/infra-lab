package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	r := gin.Default()
	setupPingRouter(r)
	setupHealthzRouter(r)
	setupPrometheusRouter(r)
	setupComputeRouter(r)
	setupLoadRouter(r)
	return r
}

func TestPingEndpoint(t *testing.T) {
	r := setupTestRouter()

	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	want := http.StatusOK
	got := w.Code
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

	var body map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatalf("failed to unmarshal body: %v", err)
	}

	wantBody := "pong"
	gotBody := body["message"]
	if wantBody != gotBody {
		t.Errorf("want %s, got %s", wantBody, gotBody)
	}
}

func TestHealthzEndpoint(t *testing.T) {
	r := setupTestRouter()

	req, _ := http.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	want := http.StatusOK
	got := w.Code
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

	var body map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatalf("failed to unmarshal body: %v", err)
	}
}

func TestPrometheusEndpoint(t *testing.T) {
	r := setupTestRouter()

	req, _ := http.NewRequest("GET", "/metrics", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	want := http.StatusOK
	got := w.Code
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

	wantBody := "http_requests_total"
	gotBody := w.Body.String()
	if !strings.Contains(gotBody, wantBody) {
		t.Errorf("want %s, got %s", wantBody, gotBody)
	}
}

func TestComputeEndpoint(t *testing.T) {
	r := setupTestRouter()

	req, _ := http.NewRequest("GET", "/compute", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	want := http.StatusOK
	got := w.Code
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

	var body map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatalf("failed to unmarshal body: %v", err)
	}

	wantBody := "ok"
	gotBody := body["message"]
	if wantBody != gotBody {
		t.Errorf("want %s, got %s", wantBody, gotBody)
	}
}

func TestLoadEndpoint(t *testing.T) {
	r := setupTestRouter()

	req, _ := http.NewRequest("GET", "/load", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	want := http.StatusOK
	got := w.Code
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

	var body map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatalf("failed to unmarshal body: %v", err)
	}

	wantBody := "ok"
	gotBody := body["message"]
	if wantBody != gotBody {
		t.Errorf("want %s, got %s", wantBody, gotBody)
	}
}