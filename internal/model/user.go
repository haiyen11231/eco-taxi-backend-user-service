package model

type User struct {
	Id                int64   `json:"id" gorm:"column:id; primaryKey; autoIncrement"`
	Name              string  `json:"name" gorm:"column:name; type:varchar(50)"`
	PhoneNumber       string  `json:"phone_number" gorm:"column:phone_number; type:varchar(8);unique"`
	Email             string  `json:"email" gorm:"column:email; type:varchar(50);unique"`
	Password          string  `json:"password" gorm:"column:password; type:varchar(255)"`
	DistanceTravelled float64 `json:"distance_travelled" gorm:"column:distance_travelled;default:0"`
}

func (User) TableName() string {
	return "users"
}

// for signup and update
type UserData struct {
	Name        string `json:"name" gorm:"column:name; type:varchar(50)"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number; type:varchar(8)"`
	Email       string `json:"email" gorm:"column:email; type:varchar(50)"`
	Password    string `json:"password" gorm:"column:password; type:varchar(255)"`
}

func (UserData) TableName() string {
	return User{}.TableName()
}

type LogInUserData struct {
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number; type:varchar(8)"`
	Password    string `json:"password" gorm:"column:password; type:varchar(255)"`
}

func (LogInUserData) TableName() string {
	return User{}.TableName()
}