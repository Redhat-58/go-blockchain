package db

import (
	"fmt"

	"go.etcd.io/bbolt"
)

var db *bbolt.DB

func InitDB() {
	newDB, err := bbolt.Open("blockchain.db", 0600, nil)
	if err != nil {
		fmt.Println("Error connecting to Database", err)
	}
	db = newDB
}
