package infrastructure

import (
	"github.com/4afS/cli-clean-architecture/domain/model"
	"github.com/4afS/cli-clean-architecture/usecase"

	"fmt"

	bolt "go.etcd.io/bbolt"
)

type UserRepository struct {
	*DbHandler
}

func NewUserRepository(dbHandler *DbHandler) usecase.UserRepository {
	userRepository := UserRepository{dbHandler}
	return &userRepository
}

func (userRepo *UserRepository) Add(user model.User) error {
	defer userRepo.DbHandler.Db.Close()

	err := userRepo.DbHandler.Db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("user"))
		if err != nil {
			return err
		}

		err = b.Put([]byte(user.Name), []byte(user.Email))

		return err
	})

	return err
}

func (userRepo *UserRepository) Remove(username string) error {
	defer userRepo.DbHandler.Db.Close()

	err := userRepo.DbHandler.Db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("user"))
		if err != nil {
			return err
		}

		err = b.Delete([]byte(username))

		return err
	})

	return err
}

func (userRepo *UserRepository) FindByName(name string) (model.User, error) {
	defer userRepo.DbHandler.Db.Close()
	var user model.User

	err := userRepo.DbHandler.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("user"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}

		email := b.Get([]byte(name))
		if email == nil {
			return fmt.Errorf("username not found")
		}

		user.Name = name
		user.Email = string(email)
		return nil
	})

	return user, err
}
