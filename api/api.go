package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Response Code, Usually 200
	Code int

	// Account Balance
	Balance int64
}

// Coin Error Response
type CoinErrorResponse struct {
	// Error Code
	Code int

	// Error Message
	Message string
}

func WriteError(w http.ResponseWriter, message string, code int) {
	resp := CoinErrorResponse{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		WriteError(w, err.Error(), http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		WriteError(w, "An unexpected error occurred.", http.StatusInternalServerError)
	}
)
