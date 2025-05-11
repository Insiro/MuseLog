package repository

import (
	"gorm.io/gorm"
	"your_spotify/internal/entity"
	"your_spotify/pkg/museLogError"
)

type User struct {
	db *gorm.DB
}

func (r User) GetUser(userSearch entity.User, includeToken bool) (entity.User, error) {
	user := entity.User{}
	result := r.db.Model(user).Preload("Settings").Where(userSearch).First(&user)

	if result.Error != nil {

		return user, museLogError.New(museLogError.NOT_FOUND, "NOT_FOUND").Wrap(result.Error)

	}

	if includeToken {
		user.AccessToken = nil
		user.RefreshToken = nil
	}

	return user, result.Error

}
