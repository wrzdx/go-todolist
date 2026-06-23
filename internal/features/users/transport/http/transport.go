package users_tranport_http

import (
	"context"
	"net/http"

	"github.com/wrzdx/go-todolist/internal/core/domain"
	core_http_server "github.com/wrzdx/go-todolist/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)

}

func NewUsersHTTPHandler(userService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: userService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: http.MethodPost,
			Path: "/users",
			Handler: h.CreateUser,
		},
	}
}
