package helpers

import (
	"mathsant/web-service-fiber/model/entity"
	"mathsant/web-service-fiber/model/response"
)

func MakeUserResponse(user *entity.User) error {
	err := response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return err
}
