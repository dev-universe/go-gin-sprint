package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

type greeting struct {
	Name     string `json:"name"`
	Greeting string `json:"greeting"`
}

type errorResp struct {
	Error  string `json:"error"`
	Detail string `json:"detail,omitempty"`
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	RegisterRoutes(r)
	return r
}

func TestHealth(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, w.Code)
	}

	expected := `{"status":"ok"}`
	if w.Body.String() != expected {
		t.Fatalf("expected body %q, got %q", expected, w.Body.String())
	}
}

func TestGreet_NoNames(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/greet", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, w.Code)
	}

	var got errorResp
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to unmarshal: %v, body=%q", err, w.Body.String())
	}

	if got.Error != "no valid names" {
		t.Fatalf("expected error %q, got %q", "no valid names", got.Error)
	}
}

func TestGreet_Success(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/greet?name=Alice&name=Bob", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, w.Code)
	}

	var got []greeting
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to unmarshal: %v, body=%q", err, w.Body.String())
	}

	want := []greeting{
		{Name: "Alice", Greeting: "Hello Alice!"},
		{Name: "Bob", Greeting: "Hello Bob!"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %#v, got %#v", want, got)
	}

}

func TestGreet_NameTooLong(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/greet?name=this-name-is-way-too-long-for-the-rule", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected %d, got %d", http.StatusUnprocessableEntity, w.Code)
	}

	var got errorResp
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to unmarshal: %v, body=%q", err, w.Body.String())
	}

	if got.Error != "name too long" {
		t.Fatalf("expected error %q, got %q", "name too long", got.Error)
	}
	if got.Detail == "" {
		t.Fatalf("expected detail to be non-empty")
	}
}
