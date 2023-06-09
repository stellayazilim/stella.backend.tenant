package dto

type ProductQueryDto struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type CategoryGetQueryDto struct {
	Limit      int             `json:"limit"`
	Offset     int             `json:"offset"`
	Categories ProductQueryDto `json:"categories"`
}
