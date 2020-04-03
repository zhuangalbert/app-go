package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/zhuangalbert/app-go/api/models"
)

var users = []models.User{
	models.User{
		Username: "zhuangalbert",
		Name:     "albert",
		Role:     "user",
		Password: "abganteng",
	},
	models.User{
		Username: "John Doe",
		Name:     "john",
		Role:     "user",
		Password: "password",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
