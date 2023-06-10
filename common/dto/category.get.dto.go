package dto

type ProductQueryDto struct {
	Limit int `json:"limit"`
}

type CategoryGetQueryDto struct {
	Limit      int             `json:"limit"`
	Categories ProductQueryDto `json:"categories"`
}
