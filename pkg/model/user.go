package model

import (
	"github.com/google/uuid"
	"github.com/michee/micgram/pkg/database"
	"github.com/michee/micgram/pkg/hook"
	"github.com/michee/micgram/pkg/utils"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type User struct {
	UserId             string `gorm:"primary_key"`
	Name               string `json:"name"`
	Username           string `json:"username"`
	Email              string `gorm:"unique"`
	Password           string `gorm:"password"`
	EmailVerified      bool
	VerificationToken  string `json:"verificationToken"`
}


func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.UserId = uuid.New().String()
	return nil
}

func init() {
	database.ConnectDB()
	DB = database.GetDB()
	DB.DropTableIfExists(&User{})
	DB.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	hashedPassword, _ := utils.HashPassword(u.Password)
	token := hook.GenerateVerificationToken()
	u.Password = hashedPassword
	u.VerificationToken = token
	DB.Create(u)
	return u
}
