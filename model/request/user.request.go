package request

type UserCreateRequest struct {
	Name    string `json:"name" validate:"required,min=3,max=50"`
	Email   string `json:"email" validate:"required,email"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}
