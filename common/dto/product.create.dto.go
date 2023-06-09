package dto

type ProductCreateDto struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Explanation string            `json:"explanation"`
	Sku         string            `json:"sku"`
	Specs       map[string]string `json:"specs"`
	Tags        []string          `json:"tags"`
	Categories  []uint            `json:"categories"`
}
