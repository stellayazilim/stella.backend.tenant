package dto

type UserCreateDto struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
