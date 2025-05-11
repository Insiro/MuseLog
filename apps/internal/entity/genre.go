package entity

import "your_spotify/internal/entity/shared"

type Genre struct {
	Id        int    `gorm:"primary_key" json:"genreId"`
	GenreName string `gorm:"size:50;not null;uniqueIndex;" json:"genreName"`
}

type GenreRelation struct {
	GenreId int `json:"genreId"`
	shared.TargetRelation
}
