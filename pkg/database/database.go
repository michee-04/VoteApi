package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB(){
	dsn := "root:@tcp(127.0.0.1:3306)/micgram?charset=utf8&parseTime=True&loc=Local"

	connect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failled to connect to database:", err)
	}

	db = connect
}

func GetDB() *gorm.DB{
	return db
}

func InitBD(models ...interface{}) error{
	ConnectDB()
	db := GetDB()

	// Suppression des tables existant
	for _, model := range models {
		if err := db.Migrator().DropTable(model); err != nil {
			return err
		}
	}

	// Migration des models
	if err := db.AutoMigrate(models...); err != nil {
		return err
	}

	return nil
}