package blockchain

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
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
	// defer DB.Close()
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

func SaveToken(t *Token) error {
	return DB.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("Token"))
		var err error

		// Create the bucket if it does not exist
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("Token"))
			if err != nil {
				return fmt.Errorf("failed to create bucket: %w", err)
			}
		}
		if err != nil {
			return fmt.Errorf("failed to serialize Token: %w", err)
		}
		serializedToken, err := json.Marshal(t)
		if err != nil {
			fmt.Println("Error handling Token")
		}
		err = bucket.Put([]byte(t.Name), serializedToken)
		if err != nil {
			return fmt.Errorf("failed to save Token: %w", err)
		}
		fmt.Println("Token saved...")
		return nil
	})
}

func UpdateBalance(b *Block, t *Token) {
	err := DB.Update(func(tx *bbolt.Tx) error {
		tokenBucket := tx.Bucket([]byte("Token"))
		blockBucket := tx.Bucket([]byte("Token"))
		if tokenBucket == nil || blockBucket == nil {
			return fmt.Errorf("failed to Load bucket")
		}

		serializedBlock, err := Serialize(b)
		if err != nil {
			return fmt.Errorf("failed to serialize block: %w", err)
		}
		serializedToken, err := json.Marshal(t)
		if err != nil {
			fmt.Println("Error handling Token")
		}

		err = blockBucket.Put([]byte(b.Hash), serializedBlock)
		if err != nil {
			return fmt.Errorf("failed to save block: %w", err)
		}
		fmt.Println("Block saved...")

		err = tokenBucket.Put([]byte(t.Name), serializedToken)
		if err != nil {
			return fmt.Errorf("failed to save Token: %w", err)
		}
		fmt.Println("Token saved...")
		return nil
	})
	if err != nil {
		fmt.Println("Error Saving data", err)
	}

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
