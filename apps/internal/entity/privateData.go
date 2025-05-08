package entity

type PrivateData struct {
	JwtPrivateKey string `gorm:"primaryKey"`
}
