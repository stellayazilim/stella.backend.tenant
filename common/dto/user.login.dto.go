package dto

type UserLoginDto struct {
	Email    string `json:"email" bind:"required"`
	Password string `json:"password" bind:"required"`
}
