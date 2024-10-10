package model

type User struct {
	Id                int     `json:"id" gorm:"column:id"`
	Name              string  `json:"name" gorm:"column:name"`
	PhoneNumber       int     `json:"phone_number" gorm:"column:phone_number"`
	Email             string  `json:"email" gorm:"column:email"`
	HashedPass        string  `json:"hashed_password" gorm:"column:hashed_password"`
	DistanceTravelled float64 `json:"distance_travelled" gorm:"column:distance_travelled"`
}

func (User) TableName() string {
	return "users"
}

type UserCreation struct {
	Id          int    `json:"-" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	PhoneNumber int    `json:"phone_number" gorm:"column:phone_number"`
	Email       string `json:"email" gorm:"column:email"`
	Password    string `json:"password" gorm:"column:hashed_password"`
}

func (UserCreation) TableName() string {
	return User{}.TableName()
}

type UserUpdate struct {
	Name        string `json:"name" gorm:"column:name"`
	PhoneNumber int    `json:"phone_number" gorm:"column:phone_number"`
	Email       string `json:"email" gorm:"column:email"`
	Password    string `json:"password" gorm:"column:hashed_password"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}