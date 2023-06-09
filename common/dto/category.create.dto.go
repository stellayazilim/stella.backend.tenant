package dto

type CategoryCreateDto struct {
	Name        string  `json:"name" bind:"required"`
	Description string  `json:"description"`
	Products    []*uint `json:"products"`
}
