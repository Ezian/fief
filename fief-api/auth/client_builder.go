package auth

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
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
		log.Fatalf("Cannot open auth database : %+v", err)
	}

	// Prepare db with each required buckets
	err = db.Update(func(tx *bolt.Tx) error {
		userBucket, err := tx.CreateBucketIfNotExists([]byte(USERS_BUCKET))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("cannot create bucket '%s'", USERS_BUCKET))
		}
		_, err = userBucket.CreateBucketIfNotExists([]byte(USERS_AUTH_BUCKET))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("cannot create bucket '%s/%s'", USERS_BUCKET, USERS_AUTH_BUCKET))
		}
		_, err = userBucket.CreateBucketIfNotExists([]byte(USERS_INFOS_BUCKET))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("cannot create bucket '%s/%s'", USERS_BUCKET, USERS_INFOS_BUCKET))
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Cannot create basic structure of the Database : %+v", err)
	}
	return db
}
