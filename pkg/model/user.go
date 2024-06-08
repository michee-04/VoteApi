package model

import (
	"fmt"

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

func FindUserByToken(token string) (*User, error) {
	var user User
	if err := DB.Where("verification_token = ?", token).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func (u *User) Verify() error {
	u.EmailVerified = true
	u.VerificationToken = ""
	if err := DB.Save(u).Error; err != nil {
		return fmt.Errorf("unable to verify user")
	}
	return nil
}

func (u *User) CanLogin() bool {
	return u.EmailVerified
}