package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/michee/micgram/pkg/database"
)


type Vote struct{
	VoteId string `gorm:"primary_key;column:voteId"`
	UserId string `json:"userId"`
	ElectionId string `json:"electionId"`
	CandidatId string `json:"candidatId"`
	Date time.Time
	User User `gorm:"foreignKey:UserId" json:"-"`
	Election Candidat `gorm:"foreignKey:ElectionId" json:"-"`
	Candidat Election `gorm:"foreignKey:CandidatId" json:"-"`
}


func (v *Vote) BeforeCreate(scope *gorm.DB) error {
	v.VoteId = uuid.New().String()
	return nil
}


func init() {
	database.ConnectDB()
	DB = database.GetDB()
	DB.DropTableIfExists(&Candidat{})
	DB.AutoMigrate(&User{})
}

func (v *Vote) CreateVote() *Vote {
	DB.Create(v)
	return v
}


func GetVote() []Vote{
	var v []Vote
	DB.Find(&v)
	return v
}


func GetVoteById(Id string) (*Vote, *gorm.DB) {
	var v Vote
	db := DB.Where("voteId", Id).Find(&v)
	return &v, db
}

func DeleteVote(Id string) Vote{
	var v Vote
	DB.Where("voteId", Id).Delete(&v)
	return v
}