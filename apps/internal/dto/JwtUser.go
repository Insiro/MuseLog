package dto

import (
	"github.com/golang-jwt/jwt/v5"
	"your_spotify/internal/constant/enum"
)

type JwtUser struct {
	UserId   string        `json:"userId"`
	UserType enum.UserType `json:"userType"`
}

type JwtClaims struct {
	JwtUser
	jwt.RegisteredClaims
}
