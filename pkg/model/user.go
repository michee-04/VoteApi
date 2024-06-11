package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/michee/micgram/pkg/database"
	"github.com/michee/micgram/pkg/hook"
	"github.com/michee/micgram/pkg/utils"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type User struct {
	UserId             string `gorm:"primary_key;column:userid"`
	Name               string `json:"name"`
	Username           string `json:"username"`
	Email              string `gorm:"unique"`
	Password           string `gorm:"password"`
	EmailVerified      bool
	VerificationToken  string `json:"verificationToken"`
	ResetToken         string `json:"resetToken"`
	ResetTokenExpiry   time.Time `json:"resetTokenExpiry"`
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

func GetAllUser() []User {
	var Users []User
	DB.Find(&Users)
	return Users
}

func GetUserById(Id string) (*User, *gorm.DB) {
	var GetUser User
	db := DB.Where("userId=?", Id).Find(&GetUser)
	return &GetUser, db
}

func DeleteUserId(Id string) User {
	var user User
	DB.Where("userId=?", Id).Delete(user)
	return user
}

// Ajout des nouvelles méthodes pour la réinitialisation et la modification du mot de passe
func (u *User) GenerateResetToken() error {
	token := uuid.New().String()
	u.ResetToken = token
	u.ResetTokenExpiry = time.Now().Add(1 * time.Hour) 
	return DB.Save(u).Error
}

func FindUserByResetToken(token string) (*User, error) {
	var user User
	if err := DB.Where("reset_token = ? AND reset_token_expiry > ?", token, time.Now()).First(&user).Error; err != nil {
		return nil, fmt.Errorf("invalid or expired reset token")
	}
	return &user, nil
}

func (u *User) UpdatePassword(newPassword string) error {
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	u.ResetToken = ""
	u.ResetTokenExpiry = time.Time{}
	
	fmt.Println("New hashed password:", hashedPassword)
    
	return DB.Save(u).Error}
