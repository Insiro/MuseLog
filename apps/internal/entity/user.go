package entity

import "time"
import "your_spotify/internal/constant/enum"

type User struct {
	Id              int        `gorm:"primaryKey:autoIncrement"`
	UserId          string     `gorm:"uniqueIndex;size:200" json:"userId"`
	UserName        string     `gorm:"size:200" json:"userName"`
	Admin           bool       `gorm:"default:false" json:"admin"`
	SpotifyId       *string    `gorm:"size:200;index" json:"spotifyId"`
	ExpireIn        int        `gorm:"type:int;default:0" json:"expireIn"`
	AccessToken     *string    `gorm:"size:200;default:null" json:"accessToken"`
	RefreshToken    *string    `gorm:"size:200;default:null" json:"refreshToken"`
	LastTimestamp   int64      `gorm:"default:0" json:"lastTimestamp"`
	LastImport      *string    `gorm:"default:null" json:"lastImport"`
	PublicToken     *string    `gorm:"size:200;default:null;index" json:"publicToken"`
	FirstListenedAt *time.Time `gorm:"type:datetime;default:null" json:"firstListenedAt"`
	//tracks: Schema.Types.ObjectId[];
	Settings *UserSetting `json:"settings"`
}

type UserSetting struct {
	UserId               int               `gorm:"primaryKey"  json:"userId"`
	HistoryLine          bool              `default:"true" gorm:"default:true" json:"historyLine"`
	PreferredStatsPeriod string            `default:"day" gorm:"default:'day'" json:"preferredStatsPeriod"`
	NbElements           int               `default:"10" gorm:"default:10" json:"nbElements"`
	MetricUsed           enum.MetricType   `default:"number" gorm:"default:'number';type:enum('number', 'duration')" json:"metricUsed"`
	DarkMode             enum.DarkModeType `default:"follow" gorm:"default:'follow';type:enum('follow', 'dark', 'light')" json:"darkMode"`
	Timezone             time.Location     `gorm:"size(10)" json:"timezone"`
	DateFormat           string            `gorm:"size(50)" json:"dateFormat"`
	BlackListedArtists   []string          `gorm:"-"`
}
