package model

import (

	"github.com/michee/micgram/pkg/database"
	"github.com/michee/micgram/pkg/utils"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct{
	UserId 	string	`gorm:"primary_key"`
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `gorm:"unique"`
	Password string `gorm:"password"`
	EmailVerified bool 
	VerificationToken string `json:"verificationToken"`
}

func (user *User) IdHexa(scope *gorm.DB){
	utils.BeforeCreateIdHexa(scope, "UserId")
}

func init(){
	if err := database.InitBD(User{}); err != nil {
		panic (err)
	}
}

func (u *User) CreateUser() *User{

	DB.Create(u)
	return u
}