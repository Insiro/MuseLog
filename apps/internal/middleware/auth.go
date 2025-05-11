package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"your_spotify/internal/app/bean"
	"your_spotify/internal/constant/enum"
	"your_spotify/internal/dto"
	"your_spotify/internal/entity"
	"your_spotify/pkg/museLogError"
)

func ValidateToken(c *gin.Context) (dto.JwtUser, error) {
	user := BaseLogged(c, true)
	if user == nil {

		return dto.JwtUser{}, museLogError.New(museLogError.NOT_LOGGED, "NOT_LOGGED")
	}

	return *user, nil
}

func parseToken(token string, jwtSecret string) *dto.JwtClaims {
	payload := dto.JwtClaims{}
	jwtUser, e := jwt.ParseWithClaims(token, &payload, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if e == nil || jwtUser == nil || !jwtUser.Valid {
		return nil
	}

	claims, ok := jwtUser.Claims.(dto.JwtClaims)
	if !ok {
		return nil
	}

	return &claims
}

func BaseLogged(c *gin.Context, useQueryToken bool) *dto.JwtUser {
	queryToken := c.Param("token")

	repo := bean.GetBean().GetRepository()
	if useQueryToken && queryToken != "" {
		user, e := repo.User().GetUser(entity.User{PublicToken: &queryToken}, false)
		if e == nil {
			userLevel := enum.User
			if user.Admin {
				userLevel = enum.Admin
			}

			return &dto.JwtUser{UserId: user.Id, UserType: userLevel}
		}
	}

	token, e := c.Cookie("token")
	if e != nil {
		return nil
	}

	privateData, e := repo.PrivateData().GetPrivateData()
	if e != nil {
		return nil
	}

	payload := parseToken(token, privateData.JwtPrivateKey)

	if payload == nil {
		return nil
	}
	return &payload.JwtUser

}
