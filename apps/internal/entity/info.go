package entity

import "time"

type InfosBlackListBy string

const (
	InfosBlackListByArtist InfosBlackListBy = "artist"
)

type Infos struct {
	Id              int
	UserId          int
	AlbumId         int
	PrimaryArtistId int
	DurationMs      int
	PlayedAt        time.Time         `gorm:"index"`
	BlackListBy     *InfosBlackListBy `gorm:"enum('artist')"`
}

type InfoArtist struct {
	InfosId  int
	ArtistId int
}
