package db

import (
	"fmt"
	"os"

	"github.com/hoonn9/nomadcoin/utils"
	bolt "go.etcd.io/bbolt"
)

type DB struct{}

func (DB) FindBlock(hash string) []byte {
	return findBlock(hash)
}
func (DB) LoadChain() []byte {
	return loadChain()
}
func (DB) SaveBlock(hash string, data []byte) {
	saveBlock(hash, data)
}
func (DB) SaveChain(data []byte) {	
	saveChain(data)
}
func (DB) DeleteAllBlocks() {
	emptyBlocks()
}


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
func InitDB() {
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
}

// key = hash 
func saveBlock(hash string, data []byte) {
	err := db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func saveChain(data []byte) {
	err := db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		return bucket.Put([]byte(checkpoint), data)
	})
	utils.HandleErr(err)
}

func loadChain() []byte {
	var data []byte

	// read-only transaction
	err := db.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})

	utils.HandleErr(err)
	
	return data
}

func Close() {
	db.Close()
}

func findBlock(hash string) []byte {
	var data []byte
	
	db.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})

	return data
}

func emptyBlocks() {
	db.Update(func(t *bolt.Tx) error {
		utils.HandleErr(t.DeleteBucket([]byte(blocksBucket)))
		_, err := t.CreateBucket([]byte(blocksBucket))
		utils.HandleErr(err)
		
		return nil
	})
}