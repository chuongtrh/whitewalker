package handler

import (
	"log"

	"github.com/tpphu/whitewalker/model"
	"github.com/tpphu/whitewalker/repo"
)

type userHandlerImpl struct {
	userRepo repo.UserRepo
	log      *log.Logger
}

func (n userHandlerImpl) get(id uint) (*model.User, Error) {
	user, err := n.userRepo.Find(id)
	if err != nil {
		return user, NewNotFoundErr(err)
	}
	return user, nil
}
