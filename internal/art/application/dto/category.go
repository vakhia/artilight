package dto

type CreateCategoryRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CategoryResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
