package database

type User struct {
	*BaseModel
	Username string `gorm:"size:255;not null;unique" binding:"required" json:"username"`
	Password string `gorm:"not null;" binding:"required" json:"password"`
	Email    string `gorm:"size:255;not null;unique" binding:"required" json:"email"`
	ApiKey  string `gorm:"not null;" json:"api_key"`
	IsAdmin  bool   `gorm:"default:false" json:"is_admin"`
}

func NewUser() *User{
	return &User{}
}
