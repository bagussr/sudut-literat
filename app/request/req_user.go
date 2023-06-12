package request

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email" example:"email@test.com"`
	Password string `json:"password" validate:"required"`
	Confirm  string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}