package models

type User struct {
	IDUser      uint   `gorm:"primaryKey;column:id_user"`
	Username    string `gorm:"unique"`
	Password    string
	NamaLengkap string
	Level       string
}

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
