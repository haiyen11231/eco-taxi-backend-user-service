package model

type User struct {
	Id                uint64  `json:"id" gorm:"column:id; primaryKey; autoIncrement"`
	Name              string  `json:"name" gorm:"column:name; type:varchar(50)"`
	PhoneNumber       string  `json:"phone_number" gorm:"column:phone_number; type:varchar(8);unique"`
	Email             string  `json:"email" gorm:"column:email; type:varchar(50);unique"`
	Password          string  `json:"password" gorm:"column:password; type:varchar(255)"`
	DistanceTravelled float64 `json:"distance_travelled" gorm:"column:distance_travelled;default:0"`
}

func (User) TableName() string {
	return "users"
}

type SignUpUserData struct {
	Name        string `json:"name" gorm:"column:name; type:varchar(50)"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number; type:varchar(8)"`
	Email       string `json:"email" gorm:"column:email; type:varchar(50)"`
	Password    string `json:"password" gorm:"column:password; type:varchar(255)"`
}

func (SignUpUserData) TableName() string {
	return User{}.TableName()
}

type LogInUserData struct {
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number; type:varchar(8)"`
	Password    string `json:"password" gorm:"column:password; type:varchar(255)"`
}

func (LogInUserData) TableName() string {
	return User{}.TableName()
}

type UpdateUserData struct {
	Name        string `json:"name" gorm:"column:name; type:varchar(50)"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number; type:varchar(8)"`
	Email       string `json:"email" gorm:"column:email; type:varchar(50)"`
}

func (UpdateUserData) TableName() string {
	return User{}.TableName()
}

type ChangePasswordUserData struct {
	NewPassword string `json:"new_password" gorm:"column:password; type:varchar(255)"`
}

func (ChangePasswordUserData) TableName() string {
	return User{}.TableName()
}

type UpdateDistanceUserData struct {
	Distance float64 `json:"distance" gorm:"column:distance_travelled"`
}

func (UpdateDistanceUserData) TableName() string {
	return User{}.TableName()
}