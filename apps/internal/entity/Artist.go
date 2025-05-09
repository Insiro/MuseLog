package entity

type Artist struct {
	Id           int     `gorm:"primaryKey" json:"id"`
	ExternalUrls []any   //TODO
	Followers    any     //TODO
	Genres       []Genre `gorm:"-" json:"genres"`
	Href         string  `gorm:"size:255"`
	Images       []Image `gorm:"-" json:"images"`
	Name         string  `gorm:"size:255" json:"name"`
	Popularity   int     `gorm:"default:0" json:"popularity"`
	Type         string  `gorm:"size:255" json:"type"`
	Uri          string  `gorm:"size:255" json:"uri"`
}
