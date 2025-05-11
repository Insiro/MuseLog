package entity

type Album struct {
	Id                   int            `json:"id"`
	AlbumType            string         `json:"album_type"`
	Artists              []Artist       `gorm:"-" json:"artists"`
	AvailableMarkets     []AlbumMarkets `json:"available_markets"`
	CopyRights           []any          `gorm:"-" json:"copy_rights"`
	ExternalIds          []any          `gorm:"-" json:"external_ids"`
	ExternalUrls         []any          `gorm:"-" json:"external_urls"`
	Genres               []Genre        `gorm:"-" json:"genres"` // out : string[]
	Href                 string         `gorm:"size:255" json:"href"`
	Images               []Image        `gorm:"-" json:"images"`
	Name                 string         `gorm:"size:255" json:"name"`
	Popularity           int            `gorm:"default:0" json:"popularity"`
	ReleaseDate          string         `gorm:"size:255" json:"release_date"`
	ReleaseDatePrecision string         `json:"release_date_precision"`
	Type                 string         `gorm:"size:50"`
	Uri                  string         `gorm:"size:255" json:"uri"`
}

type AlbumMarkets struct {
	AlbumId int      `json:"album_id"`
	Markets []string `gorm:"size:20" json:"markets"`
}

type SpotifyAlbum struct {
	Album
	Artists []Artist `json:"artists"`
	Track   []Track  `json:"track"`
}
