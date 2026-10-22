package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sebas3261/calculator/backend/internal/calculator"
	"github.com/sebas3261/calculator/backend/internal/calculator/evaluator"
	"github.com/sebas3261/calculator/backend/internal/calculator/parser"
	apphttp "github.com/sebas3261/calculator/backend/internal/http"
)

func newTestRouter() http.Handler {
	service := calculator.NewService(parser.NewExpressionParser(), evaluator.NewEvaluator())
	return apphttp.NewRouter(service)
}

func doCalculate(t *testing.T, router http.Handler, body string) *httptest.ResponseRecorder {
	t.Helper()

	req := httptest.NewRequest(http.MethodPost, "/calculate", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	return rec
}

func TestCalculateHandler_ValidExpression(t *testing.T) {
	router := newTestRouter()

	rec := doCalculate(t, router, `{"expression":"2 + 3 * 4"}`)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", rec.Code, http.StatusOK, rec.Body.String())
	}

	var resp calculator.CalculationResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if resp.Answer != 14 {
		t.Errorf("answer = %v, want 14", resp.Answer)
	}
}

func TestCalculateHandler_DivisionByZero(t *testing.T) {
	router := newTestRouter()

	rec := doCalculate(t, router, `{"expression":"1 / 0"}`)

	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("status = %d, want %d; body = %s", rec.Code, http.StatusUnprocessableEntity, rec.Body.String())
	}
}

func TestCalculateHandler_InvalidJSON(t *testing.T) {
	router := newTestRouter()

	rec := doCalculate(t, router, `{"expression": `)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d; body = %s", rec.Code, http.StatusBadRequest, rec.Body.String())
	}
}

func TestCalculateHandler_EmptyExpression(t *testing.T) {
	router := newTestRouter()

	rec := doCalculate(t, router, `{"expression":""}`)

	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("status = %d, want %d; body = %s", rec.Code, http.StatusUnprocessableEntity, rec.Body.String())
	}

	var resp map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if resp["error"] == "" {
		t.Error("expected a non-empty error message")
	}
}

func TestCalculateHandler_MethodNotAllowed(t *testing.T) {
	router := newTestRouter()

	req := httptest.NewRequest(http.MethodGet, "/calculate", bytes.NewReader(nil))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d; body = %s", rec.Code, http.StatusMethodNotAllowed, rec.Body.String())
	}
}

func TestHealthHandler(t *testing.T) {
	router := newTestRouter()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s", rec.Code, http.StatusOK, rec.Body.String())
	}
}
