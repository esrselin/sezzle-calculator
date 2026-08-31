package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type CalculateRequest struct {
	FirstNumber  float64 `json:"firstNumber"`
	Operation    string  `json:"operation"`
	SecondNumber float64 `json:"secondNumber"`
}

type CalculateResponse struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func calculate(first float64, operation string, second float64) (float64, error) {
	if operation == "+" {
		return first + second, nil
	}

	if operation == "-" {
		return first - second, nil
	}

	if operation == "*" {
		return first * second, nil
	}

	if operation == "/" {
		if second == 0 {
			return 0, fmt.Errorf("Cannot divide by zero.")
		}

		return first / second, nil
	}

	if operation == "^" {
		return math.Pow(first, second), nil
	}

	if operation == "sqrt" {
		if first < 0 {
			return 0, fmt.Errorf("Cannot calculate square root of a negative number.")
		}

		return math.Sqrt(first), nil
	}

	if operation == "%" {
		return first * second / 100, nil
	}

	return 0, fmt.Errorf("invalid operation")
}

func calculator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request CalculateRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := calculate(
		request.FirstNumber,
		request.Operation,
		request.SecondNumber,
	)

	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(CalculateResponse{
		Result: result,
	})
}

func writeError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(ErrorResponse{
		Error: message,
	})
}

func main() {
	http.HandleFunc("/calculate", calculator)

	fmt.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}