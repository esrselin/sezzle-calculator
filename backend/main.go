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

func calculator(w http.ResponseWriter, r *http.Request) {
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

	var result float64

	if request.Operation == "+" {
		result = request.FirstNumber + request.SecondNumber

	} else if request.Operation == "-" {
		result = request.FirstNumber - request.SecondNumber

	} else if request.Operation == "*" {
		result = request.FirstNumber * request.SecondNumber

	} else if request.Operation == "/" {
		if request.SecondNumber == 0 {
			writeError(w, "Cannot divide by zero", http.StatusBadRequest)
			return
		}

		result = request.FirstNumber / request.SecondNumber

	} else if request.Operation == "^" {
		result = math.Pow(request.FirstNumber, request.SecondNumber)

	} else {
		writeError(w, "Bad Request", http.StatusBadRequest)
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