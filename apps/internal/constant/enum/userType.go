package enum

type UserType string

const (
	Admin UserType = "ADMIN"
	User  UserType = "USER"
	Guest UserType = "GUEST"
)
