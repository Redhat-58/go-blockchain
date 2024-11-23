package controllers

import (
	"encoding/json"
	"go-blockchain/types"
	"io"
	"net/http"
)

func CreateBlock(w http.ResponseWriter, r *http.Request) {
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

	var data types.NewBlock
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Ivalid Json format", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message": "Data received",
		"data":    data.Data,
	}
	json.NewEncoder(w).Encode(response)
}
