package controllers

import (
	"encoding/json"
	"go-blockchain/types"
	"io"
	"net/http"
)

func TransferToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invald request method", http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request bodys", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var data types.NewTranasaction
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Ivalid Json format", http.StatusBadRequest)
		return
	}
	result, err := BC.Token.Transfer(BC, data.From, data.To, data.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": result,
		"from":    data.From,
		"to":      data.To,
		"amount":  data.Amount,
	}
	json.NewEncoder(w).Encode(response)
}
