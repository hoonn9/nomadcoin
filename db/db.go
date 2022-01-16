package db

import (
	"fmt"
	"os"

	"github.com/hoonn9/nomadcoin/utils"
	bolt "go.etcd.io/bbolt"
)

const (
	dbName = "blockchain"
	dataBucket = "data"
	blocksBucket = "blocks"

	checkpoint = "checkpoint"
)

var db *bolt.DB


func getDbName() string {
	port := os.Args[2][6:]
	return fmt.Sprintf("%s_%s.db", dbName, port)
}

// singleton
func DB() *bolt.DB {
	if db == nil {
		// init db
		dbPointer, err := bolt.Open(getDbName(), 0600, nil)
		utils.HandleErr(err)
		db = dbPointer

		// Bucket(sql의 table 같은) 생성
		// transaction (read, write only byte)
		err = db.Update(func(t *bolt.Tx) error {
			_, err = t.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)
	}

	return db
}

// key = hash 
func SaveBlock(hash string, data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		return bucket.Put([]byte(hash), data)
	})
	utils.HandleErr(err)
}

func SaveCheckpoint(data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		return bucket.Put([]byte(checkpoint), data)
	})
	utils.HandleErr(err)
}

func Checkpoint() []byte {
	var data []byte

	// read-only transaction
	err := DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})

	utils.HandleErr(err)
	
	return data
}

func Close() {
	DB().Close()
}

func Block(hash string) []byte {
	var data []byte
	
	DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})

	return data
}