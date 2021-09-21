package auth

import (
	"encoding/json"
	"fief/models"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// UserInfos User informations
type UserInfos struct {
	// email
	Email string `json:"email"`

	// login
	Login string `json:"login"`
}

func RegisterNewUser(db *bolt.DB, userInfo *models.RegisterUser) error {
	err := db.Update(func(tx *bolt.Tx) error {
		usersBucket := tx.Bucket([]byte(USERS_BUCKET))
		if usersBucket == nil {
			return fmt.Errorf("no bucket %q", USERS_BUCKET)
		}
		err := createUserAuth(usersBucket, userInfo)
		if err != nil {
			return errors.Wrapf(err, "cannot create user %q", *userInfo.Login)
		}
		err = createUserInfos(usersBucket, userInfo)
		if err != nil {
			return errors.Wrapf(err, "cannot create user %q", *userInfo.Login)
		}
		return nil
	})
	return err
}

func createUserAuth(usersBucket *bolt.Bucket, userInfo *models.RegisterUser) error {
	userKey := []byte(*userInfo.Login)
	password := []byte(*userInfo.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	authBucket, err := safeGetUsersSubBucket(usersBucket, USERS_AUTH_BUCKET, userKey)
	if err != nil {
		return err
	}
	authBucket.Put(userKey, hashedPassword)
	return nil
}

func createUserInfos(usersBucket *bolt.Bucket, userInfo *models.RegisterUser) error {
	userKey := []byte(*userInfo.Login)
	infosBucket, err := safeGetUsersSubBucket(usersBucket, USERS_INFOS_BUCKET, userKey)
	if err != nil {
		return err
	}

	infos := UserInfos{
		Login: *userInfo.Login,
		Email: *userInfo.Email,
	}

	encoded, err := json.Marshal(infos)
	if err != nil {
		return errors.Wrapf(err, "cannot encode information from user %q", *userInfo.Login)
	}
	infosBucket.Put(userKey, encoded)
	return nil
}

func safeGetUsersSubBucket(bucket *bolt.Bucket, subKey string, userKey []byte) (*bolt.Bucket, error) {
	subBucket := bucket.Bucket([]byte(subKey))
	if subBucket == nil {
		return nil, fmt.Errorf("no bucket %q in bucket %q", subKey, USERS_BUCKET)
	}
	if subBucket.Get(userKey) != nil {
		return nil, fmt.Errorf("user with login %q already registered", string(userKey))
	}
	return subBucket, nil
}

func CheckUserPassword(db *bolt.DB, login, password string) (bool, error) {

	// TODO get password from DB
	var dbPassword []byte
	err := db.View(func(tx *bolt.Tx) error {
		usersBucket := tx.Bucket([]byte(USERS_BUCKET))
		if usersBucket == nil {
			return fmt.Errorf("no bucket %q", USERS_BUCKET)
		}
		authBucket := usersBucket.Bucket([]byte(USERS_AUTH_BUCKET))
		if usersBucket == nil {
			return fmt.Errorf("no bucket %q", USERS_AUTH_BUCKET)
		}

		dbPassword = authBucket.Get([]byte(login))
		return nil
	})
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(dbPassword, []byte(password))
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	return true, nil
}
