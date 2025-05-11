package bean

import "your_spotify/internal/repository"

type Repository struct {
	user        repository.User
	privateData repository.PrivateData
}

func (r Repository) User() repository.User {
	return r.user
}

func (r Repository) PrivateData() repository.PrivateData {
	return r.privateData
}
