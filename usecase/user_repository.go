package usecase

import (
	"github.com/4afS/cli-clean-architecture/domain/model"
)

type UserRepository interface {
	Add(model.User) error
	Remove(string) error
	FindByName(string) (model.User, error)
}
