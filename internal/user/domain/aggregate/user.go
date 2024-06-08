package aggregate

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	FirstName   string
	LastName    string
	Avatar      string
	Cover       string
	Email       string
	Bio         string
	Position    string
	Gender      string
	Currency    string
	PhoneNumber string
	Location    string
	Password    string
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewUser(firstName, lastName, email, password string) (User, error) {
	return User{
		Id:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}, nil
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetAvatar(path string) {
	u.Avatar = path
}

func (u *User) SetCover(path string) {
	u.Cover = path
}

func (u *User) SetPosition(position string) {
	u.Position = position
}

func (u *User) SetGender(gender string) {
	u.Gender = gender
}

func (u *User) SetCurrency(currency string) {
	u.Currency = currency
}

func (u *User) SetPhoneNumber(phoneNumber string) {
	u.PhoneNumber = phoneNumber
}

func (u *User) SetLocation(location string) {
	u.Location = location
}

func (u *User) SetAddress(address string) {
	u.Address = address
}

func (u *User) SetBio(bio string) {
	u.Bio = bio
}
