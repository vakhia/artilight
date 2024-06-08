package entity

type Category struct {
	Id          int
	Title       string
	Description string
}

func NewCategory(name, description string) Category {
	return Category{
		Title:       name,
		Description: description,
	}
}
