package db

import (
	"github.com/boltdb/bolt"
	"github.com/hoonn9/nomadcoin/utils"
)

const (
	dbName = "blockchain.db"
	dataBucket = "data"
	blocksBucket = "blocks"
)

var db *bolt.DB

// bucket (sql의 table 같은)

// singleton
func DB() *bolt.DB {
	if db == nil {
		// init db
		dbPointer, err := bolt.Open("blockchain.db", 0600, nil)
		utils.HandleErr(err)
		
		// Bucket 생성
		// transaction (read, write)
		err = db.Update(func(t *bolt.Tx) error {
			_, err = t.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)
		db = dbPointer
	}

	return db
}