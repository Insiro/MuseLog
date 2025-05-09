package entity

import "your_spotify/internal/entity/shared"

type Image struct {
	Id int `gorm:"primary_key" json:"id"`
	shared.TargetRelation
	Url    string `gorm:"size:255" json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
