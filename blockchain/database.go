package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"go.etcd.io/bbolt"
)

var DB *bbolt.DB

func InitDB() {
	newDB, err := bbolt.Open("blockchain.db", 0666, nil)
	if err != nil {
		fmt.Println("Error connecting to Database", err)
	}
	DB = newDB
	defer DB.Close()
}

func SaveBlock(b *Block) error {
	return DB.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("blocks"))
		var err error

		// Create the bucket if it does not exist
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("blocks"))
			if err != nil {
				return fmt.Errorf("failed to create bucket: %w", err)
			}
		}
		serializedBlock, err := Serialize(b)
		if err != nil {
			return fmt.Errorf("failed to serialize block: %w", err)
		}

		err = bucket.Put([]byte(b.Hash), serializedBlock)
		if err != nil {
			return fmt.Errorf("failed to save block: %w", err)
		}
		fmt.Println("Block saved...")
		return nil
	})
}

func Serialize(block *Block) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize block: %w", err)
	}
	return buffer.Bytes(), nil
}
