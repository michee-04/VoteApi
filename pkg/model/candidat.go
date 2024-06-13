package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/michee/micgram/pkg/database"
)


type Candidat struct{
	CandidatId string `gorm:"primary_key;column:candidatId"`
	Name string `json:"name"`
	Email string `gorm:"unique"`
	Discour string `json:"discour"`
	ElectionId string `json:"electionId"`
	Election Election `gorm:"foreignKey:ElectionId" json:"-"`
	Vote []Vote `gorm:"foreignKey:CandidatId"`
}

func (candidat *Candidat) BeforeCreate(scope *gorm.DB) error {
	candidat.CandidatId = uuid.New().String()
	return nil
}


func init() {
	database.ConnectDB()
	DB = database.GetDB()
	DB.DropTableIfExists(&Candidat{})
	DB.AutoMigrate(&Candidat{})
}

func (c *Candidat) CreateCandidat() *Candidat {
	DB.Create(c)
	return c
}


func GetCandidat() []Candidat{
	var c []Candidat
	DB.Preload("Election").Preload("Vote").Find(&c)
	return c
}


func GetCandidatById(Id string) (*Candidat, *gorm.DB){
	var c Candidat
	db := DB.Preload("Election").Preload("Vote").Where("candidatId", Id).Find(&c)

	return &c, db
}

func DeleteCandidat(Id string) Candidat{
	var c Candidat
	DB.Where("candidatId=?", Id).Delete(c)
	return c
}