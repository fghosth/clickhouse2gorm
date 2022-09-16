package model

import "time"

type Issues struct {
	Name       string    `db:"name"`
	Owner      string    `db:"owner"`
	Category   string    `db:"category"`
	Epics      string    `db:"epics"`
	Project    string    `db:"project"`
	DueDate    time.Time `db:"dueDate"`   //nolint:typecheck
	StartDate  time.Time `db:"startDate"` //nolint:typecheck
	Status     string    `db:"status"`
	EndDate    time.Time `db:"endDate"` //nolint:typecheck
	Progress   int16     `db:"progress"`
	Key        string    `db:"key"`
	CreateDate time.Time `db:"createDate"` //nolint:typecheck
}
