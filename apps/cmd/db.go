package loader

import (
	"fmt"
	"gorm.io/gorm"
	"your_spotify/internal/entity"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&entity.GlobalPreference{}, &entity.PrivateData{}, &entity.User{}, &entity.UserSetting{}); err != nil {
		fmt.Println("Migration failed")
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println("Migration successful")

}
