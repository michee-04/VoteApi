package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/michee/micgram/pkg/database"
)

type Admin struct{
	AdminId string `gorm:"primary_key;column:adminId"`
	UserId string `json:"userId"`
	User User `gorm:"foreignKey:UserId" json:"-"`
}


func (a *Admin) BeforeCreate(scope *gorm.DB) error {
	a.AdminId = uuid.NewString()
	return nil
}


func init() {
	database.ConnectDB()
	DB = database.GetDB()
	DB.AutoMigrate(&Admin{})
}


func (a *Admin) CreateAdmin() *Admin{
	DB.Create(a)
	return a
}


func GetAdminByID(Id string) (*Admin, *gorm.DB) {
	var a Admin
	db := DB.Where("adminId=?", Id).Find(&a)
	return &a, db
}