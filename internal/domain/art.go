package domain

type Art struct {
	BaseModel
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryId  int     `json:"category_id"`
	OwnerId     int     `json:"owner_id"`
	Category    Category
	Owner       User
}
