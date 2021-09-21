package auth

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type ClientBuilder struct {
	// this place can be used to initialize the auth client which can be used to talk to other micro-services
}

func NewClientBuilder() ClientBuilder {
	return ClientBuilder{}
}

func (b ClientBuilder) BuildBoltClient() *bolt.DB {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	// TODO set auth db file as configuration/environment variable
	db, err := bolt.Open("auth.db", 0600, nil)
	if err != nil {
		log.Fatal("Cannot open auth database : ", err.Error())
	}

	// Prepare db with each required buckets
	err = db.Update(func(tx *bolt.Tx) error {
		userBucket, err := tx.CreateBucketIfNotExists([]byte(USERS_BUCKET))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		_, err = userBucket.CreateBucketIfNotExists([]byte(USERS_AUTH_BUCKET))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		_, err = userBucket.CreateBucketIfNotExists([]byte(USERS_INFOS_BUCKET))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	if err != nil {
		log.Fatal("error connecting DB : ", err.Error())
	}
	return db
}
