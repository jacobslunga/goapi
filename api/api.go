package api

import (
  "encoding/json"
  "net/http"
)

type CoinBalanceParams struct {
  Username string `json:"username"`
}

type CoinBalanceResponse struct {
  Code int `json:"code"`
  Balance int64 `json:"balance"`
}

type Error struct {
  Code int `json:"code"`
  Message string `json:"message"`
}

func writeError(w http.ResponseWriter, message string, code int) {
  resp := Error{
    Code: code,
    Message: message,
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)

  json.NewEncoder(w).Encode(resp)
}

var (
  RequestErrorHandler = func(w http.ResponseWriter, err error) {
    writeError(w, err.Error(), http.StatusBadRequest)
  }
  InternalErrorHandler = func(w http.ResponseWriter) {
    writeError(w, "An unexpected error occurred", http.StatusInternalServerError)
  }
)
