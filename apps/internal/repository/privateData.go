package repository

import (
	"gorm.io/gorm"
	"your_spotify/internal/entity"
	"your_spotify/pkg"
)

type PrivateData struct {
	db *gorm.DB
}

func (r PrivateData) CreatePrivateData() error {

	token, err := pkg.Crypt.GenerateRandomString(32)

	if err != nil {
		return err
	}

	prvData := entity.PrivateData{JwtPrivateKey: token}
	r.db.Create(&prvData)
	return nil

}

func (r PrivateData) GetPrivateData() (entity.PrivateData, error) {
	privateData := entity.PrivateData{}
	result := r.db.Last(&privateData)
	if result.Error != nil {
		return privateData, result.Error
	}

	return privateData, nil
}
