package api

import (
	"encode/json"
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
