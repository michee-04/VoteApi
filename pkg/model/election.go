package model

import (

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/michee/micgram/pkg/database"
)


type Election struct{
	ElectionId string `gorm:"primary_key;column:electionId"`
	Image string `json:"image"`
	Name string `json:"name"`
	Description string `json:"description"`
	StartDate string `json:"startDate"` 
	EndDate string `json:"endDate"`
	Candidat []Candidat `gorm:"foreignKey:ElectionId"`
	Vote []Vote `gorm:"foreignKey:ElectionId"`
}

func (election *Election) BeforeCreate(scope *gorm.DB) error {
	election.ElectionId = uuid.New().String()
	return nil
}


func init() {
	database.ConnectDB()
	DB = database.GetDB()
	// DB.DropTableIfExists(&Election{})
	DB.AutoMigrate(&Election{})
}

func (e *Election) CreateElection() *Election {
	DB.Create(e)
	return e
}

func GetAllElection() []Election {
	var getElection []Election
	DB.Preload("Candidat").Preload("Vote").Find(&getElection)
	return getElection
}

func GetElectionById(Id string) (*Election, *gorm.DB){
	var getElection Election
	db := DB.Preload("Candidat").Preload("Vote").Where("electionId", Id).Find(&getElection)

	return &getElection, db
}

func DeleteElection(Id string) Election {
	var election Election
	DB.Where("electionId=?", Id).Delete(election)
	return election
}