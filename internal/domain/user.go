package domain

type User struct {
	BaseModel
	Name     string
	Email    string
	Password string
}
