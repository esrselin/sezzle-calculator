package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name      string
		first     float64
		operation string
		second    float64
		expected  float64
		wantError bool
	}{
		{
			name:      "addition",
			first:     10,
			operation: "+",
			second:    5,
			expected:  15,
		},
		{
			name:      "subtraction",
			first:     10,
			operation: "-",
			second:    5,
			expected:  5,
		},
		{
			name:      "multiplication",
			first:     10,
			operation: "*",
			second:    5,
			expected:  50,
		},
		{
			name:      "division",
			first:     10,
			operation: "/",
			second:    5,
			expected:  2,
		},
		{
			name:      "exponentiation",
			first:     2,
			operation: "^",
			second:    3,
			expected:  8,
		},
		{
			name:      "division by zero",
			first:     10,
			operation: "/",
			second:    0,
			expected:  0,
			wantError: true,
		},
		{
			name:      "negative numbers",
			first:     -10,
			operation: "+",
			second:    5,
			expected:  -5,
		},
		{
			name:      "decimal numbers",
			first:     2.5,
			operation: "*",
			second:    4,
			expected:  10,
		},
		{
			name:      "invalid operation",
			first:     10,
			operation: "%",
			second:    5,
			expected:  0,
			wantError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := calculate(
				test.first,
				test.operation,
				test.second,
			)

			if test.wantError {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestCalculator(t *testing.T) {
	requestBody := `{
		"firstNumber": 10,
		"operation": "+",
		"secondNumber": 5
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/calculate",
		strings.NewReader(requestBody),
	)

	response := httptest.NewRecorder()

	calculator(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	expected := `{"result":15}`

	if strings.TrimSpace(response.Body.String()) != expected {
		t.Errorf(
			"expected %s, got %s",
			expected,
			strings.TrimSpace(response.Body.String()),
		)
	}
}
func TestCalculatorDivisionByZero(t *testing.T) {
	requestBody := `{
		"firstNumber": 10,
		"operation": "/",
		"secondNumber": 0
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/calculate",
		strings.NewReader(requestBody),
	)

	response := httptest.NewRecorder()

	calculator(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", response.Code)
	}

	expected := `{"error":"cannot divide by zero"}`

	if strings.TrimSpace(response.Body.String()) != expected {
		t.Errorf(
			"expected %s, got %s",
			expected,
			strings.TrimSpace(response.Body.String()),
		)
	}
}
func TestCalculatorInvalidJSON(t *testing.T) {
	requestBody := `{
		"firstNumber": 10,
		"operation": +
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/calculate",
		strings.NewReader(requestBody),
	)

	response := httptest.NewRecorder()

	calculator(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", response.Code)
	}

	expected := `{"error":"Invalid request body"}`

	if strings.TrimSpace(response.Body.String()) != expected {
		t.Errorf(
			"expected %s, got %s",
			expected,
			strings.TrimSpace(response.Body.String()),
		)
	}
}
func TestCalculatorInvalidOperation(t *testing.T) {
	requestBody := `{
		"firstNumber": 10,
		"operation": "%",
		"secondNumber": 5
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/calculate",
		strings.NewReader(requestBody),
	)

	response := httptest.NewRecorder()

	calculator(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", response.Code)
	}

	expected := `{"error":"invalid operation"}`

	if strings.TrimSpace(response.Body.String()) != expected {
		t.Errorf(
			"expected %s, got %s",
			expected,
			strings.TrimSpace(response.Body.String()),
		)
	}
}
func TestCalculatorMethodNotAllowed(t *testing.T) {
	request := httptest.NewRequest(
		http.MethodGet,
		"/calculate",
		nil,
	)

	response := httptest.NewRecorder()

	calculator(response, request)

	if response.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", response.Code)
	}

	expected := `{"error":"Method not allowed"}`

	if strings.TrimSpace(response.Body.String()) != expected {
		t.Errorf(
			"expected %s, got %s",
			expected,
			strings.TrimSpace(response.Body.String()),
		)
	}
}
func TestCalculatorOptions(t *testing.T) {
	request := httptest.NewRequest(
		http.MethodOptions,
		"/calculate",
		nil,
	)

	response := httptest.NewRecorder()

	calculator(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	if response.Header().Get("Access-Control-Allow-Origin") != "http://localhost:5173" {
		t.Errorf("unexpected Access-Control-Allow-Origin header")
	}

	if response.Header().Get("Access-Control-Allow-Methods") != "POST, OPTIONS" {
		t.Errorf("unexpected Access-Control-Allow-Methods header")
	}

	if response.Header().Get("Access-Control-Allow-Headers") != "Content-Type" {
		t.Errorf("unexpected Access-Control-Allow-Headers header")
	}
}