package entity

type GlobalPreference struct {
	AllowRegistrations bool `default:"true" gorm:"default:true" json:"allowRegistrations"`
	AllowAffinity      bool `default:"true" gorm:"default:true" json:"allowAffinity"`
}
