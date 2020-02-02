package controller

import (
	"github.com/4afS/cli-clean-architecture/usecase"
	"github.com/4afS/cli-clean-architecture/domain/model"
)

type UserController struct {
	UserInteractor usecase.UserInteractor
}

func NewUserController(repo usecase.UserRepository) *UserController {
	return &UserController{
		UserInteractor: usecase.UserInteractor{
			UserRepository: repo,
		},
	}
}

func (controller *UserController) Register(name string, email string) (err error) {
	err = controller.UserInteractor.UserRepository.Add(model.User{Name: name, Email: email})
	return
}

func (controller *UserController) Delete(name string) (err error) {
	err = controller.UserInteractor.UserRepository.Remove(name)
	return
}

func (controller *UserController) Search(name string) (user model.User, err error) {
	user, err = controller.UserInteractor.UserRepository.FindByName(name)
	return
}