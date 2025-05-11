package entity

type Track struct {
	Id               int
	Album            string   `gorm:"size:255" json:"album"`
	Artists          []Artist `gorm:"-" json:"album_artist"`
	AvailableMarkets []TrackMarkets
	DiskNumber       int `json:"disk_number"`
	DurationMs       int `json:"duration_ms"`
	Explicit         bool
	ExternalIds      any
	ExternalUrls     any
	Href             string `gorm:"size:255" json:"href"`
	IsLocal          bool   `json:"is_local"`
	Name             string `gorm:"size:255" json:"name"`
	Popularity       int    `json:"popularity"`
	PreviewUrl       string `gorm:"size:255" json:"preview_url"`
	TrackNumber      int    `json:"track_number"`
	Type             string `gorm:"size:255" json:"type"`
	Uri              string `gorm:"size:255" json:"uri"`
}

type TrackMarkets struct {
	TrackId int      `json:"track_id"`
	Markets []string `gorm:"size:20" json:"markets"`
}

type TrackArtists struct {
	TrackId  int `gorm:"uniqueIndex:track_artists_idx"`
	ArtistId int `gorm:"uniqueIndex:track_artists_idx;index"`
}

type SpotifyTrack struct {
	Artists []Artist
	Album   SpotifyAlbum
}

type RecentlyPlayedTrack struct {
	PlayedAt string `json:"played_at"`
	track    SpotifyTrack
}
