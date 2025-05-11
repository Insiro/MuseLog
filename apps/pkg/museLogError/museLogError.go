package museLogError

import "fmt"

type MuseLogError struct {
	error error
	Type  Type
	Msg   string
}

func New(errType Type, Msg string) *MuseLogError {
	return &MuseLogError{Type: errType, Msg: Msg}
}

func (m *MuseLogError) Wrap(err error) *MuseLogError {
	m.error = wrap(err, fmt.Sprintf("code:%d | msg:%s", m.Type, m.Msg))
	return m
}

func (m *MuseLogError) Error() string {
	return fmt.Sprintf("code:%d | msg:%s", m.Type, m.Msg)
}

func (m *MuseLogError) Unwrap() error {
	return m.error // 래핑된 에러 반환
}
