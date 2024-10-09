package model

// docker run --name eco-taxi-user-service-db-mysql -e MYSQL_ROOT_PASSWORD=user-service -d -p 3306:3306 mysql
// docker run --name eco-taxi-user-service-redis -p 6379:6379 -d redis

type User struct {
	Id                int     `json:"id,omitempty" gorm:"column:id"`
	Name              string  `json:"name,omitempty" gorm:"column:name"`
	PhoneNumber       int     `json:"phone_number" gorm:"column:phone_number"`
	Email             string  `json:"email,omitempty" gorm:"column:email"`
	HashedPass        string  `json:"hashed_password,omitempty" gorm:"column:hashed_password"`
	DistanceTravelled float64 `json:"distance_travelled,omitempty" gorm:"column:distance_travelled"`
}

func (User) TableName() string{
	return "users"
}

type UserCreation struct {
	Id                int     `json:"-" gorm:"column:id"`
	Name              string  `json:"name,omitempty" gorm:"column:name"`
	PhoneNumber       int     `json:"phone_number" gorm:"column:phone_number"`
	Email             string  `json:"email,omitempty" gorm:"column:email"`
	Password          string  `json:"password,omitempty" gorm:"column:hashed_password"`
}

func (UserCreation) TableName() string{
	return User{}.TableName()
}

type UserUpdate struct {
	Name              string  `json:"name,omitempty" gorm:"column:name"`
	PhoneNumber       int     `json:"phone_number" gorm:"column:phone_number"`
	Email             string  `json:"email,omitempty" gorm:"column:email"`
	Password          string  `json:"password,omitempty" gorm:"column:hashed_password"`
}

func (UserUpdate) TableName() string{
	return User{}.TableName()
}