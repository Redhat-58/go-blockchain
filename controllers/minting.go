package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebsocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Erro upgrading connection", err)
	}
	defer conn.Close()
	for {
		_, hash, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message", err)
			break
		}
		var data struct {
			Hash string `json:"hash"`
		}
		if err := json.Unmarshal(hash, &data); err != nil {
			fmt.Println("Invalid json format", err)
		}

		userBlock, err := BC.GetBlockByHash(data.Hash)
		if err != nil {
			fmt.Println("Error fetching block", err)
			conn.WriteMessage(websocket.TextMessage, []byte("Invalid json format"))
			continue
		}
		conn.WriteMessage(websocket.TextMessage, []byte("Mining Token...."))
		userBlock.MintToken(&BC.Token, BC)

		conn.WriteMessage(websocket.TextMessage, []byte("Token minied"))
	}
}
