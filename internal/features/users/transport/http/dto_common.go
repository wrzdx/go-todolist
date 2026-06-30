package users_tranport_http

import "github.com/wrzdx/go-todolist/internal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id"           example:"1"`
	Version     int     `json:"version"      example:"1"`
	FullName    string  `json:"full_name"    example:"Ivan Ivanov"`
	PhoneNumber *string `json:"phone_number" example:"+71112223344"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}
