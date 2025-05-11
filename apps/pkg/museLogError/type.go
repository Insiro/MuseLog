package museLogError

type Type int

const (
	DEFAULT   Type = iota
	NOT_FOUND Type = iota

	//unAuthorized
	NOT_LOGGED Type = iota
)
