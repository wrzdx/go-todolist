package users_service

import (
	"context"

	"github.com/wrzdx/go-todolist/internal/core/domain"
)

type UsersService struct {
	userRepository  UsersRepository
}

type UsersRepository interface {
	CreateUser (
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUsersService(
	usersRepository UsersRepository,
) *UsersService {
	return &UsersService{
		userRepository: usersRepository,
	}
}