package service

import (
	"demo/api/user"
	"demo/internal/user/biz"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	user.UnimplementedUserServer

	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {

	return &UserService{
		uc: uc,
	}
}
