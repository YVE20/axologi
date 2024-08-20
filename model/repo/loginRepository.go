package repo

import (
	"axologi/model/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/axologi?parseTime=true"))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Customers{})

	DB = db
}
