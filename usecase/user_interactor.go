package usecase

import (
	"github.com/4afS/cli-clean-architecture/domain/model"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Register(user model.User) (err error) {
	err = interactor.UserRepository.Add(user)
	return
}

func (interactor *UserInteractor) Delete(username string) (err error) {
	err = interactor.UserRepository.Remove(username)
	return
}

func (interactor *UserInteractor) Search(username string) (user model.User, err error) {
	user, err = interactor.UserRepository.FindByName(username)
	return
}
