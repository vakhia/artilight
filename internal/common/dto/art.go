package dto

type ArtCreateRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	CategoryId  int     `json:"category_id" binding:"required"`
	OwnerId     int     `json:"owner_id" binding:"required"`
}

type ArtUpdateRequest struct {
	Id          int
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryId  int     `json:"category_id"`
	OwnerId     int     `json:"owner_id"`
}

type ArtResponse struct {
	Id          int               `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Price       float64           `json:"price"`
	Category    *CategoryResponse `json:"category,omitempty"`
	Owner       *UserResponse     `json:"owner,omitempty"`
}
