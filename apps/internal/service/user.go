package service

import (
	"your_spotify/internal/entity"
	"your_spotify/internal/repository"
	"your_spotify/pkg/museLogError"
)

type User struct {
	userRepository repository.User
}

func (ur *User) FindUserById(userId string) (entity.User, *museLogError.MuseLogError) {
	user, e := ur.userRepository.GetUser(entity.User{Id: userId}, false)
	if e != nil {
		return user, museLogError.New(museLogError.NOT_FOUND, "NOT_FOUND").Wrap(e)
	}

	return user, nil
}
