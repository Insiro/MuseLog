package entity

import "time"

type ImportStatus string

const (
	Progress       ImportStatus = "progress"
	Success        ImportStatus = "success"
	Failure        ImportStatus = "failure"
	FailureRemoved ImportStatus = "failure-removed"
)

type ImporterState struct {
	Type      string `gorm:"size:255"`
	Total     int
	Current   int
	MetaData  *any `gorm:"-"`
	UserId    int
	Status    ImportStatus `default:"progress" gorm:"enum('progress','success', 'failure','failure-removed')'"`
	CreatedAt time.Time
}
